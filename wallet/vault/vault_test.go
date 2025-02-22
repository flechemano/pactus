package vault

import (
	"fmt"
	"strings"
	"testing"

	"github.com/pactus-project/pactus/crypto"
	"github.com/pactus-project/pactus/crypto/bls"
	"github.com/pactus-project/pactus/crypto/bls/hdkeychain"
	"github.com/pactus-project/pactus/util/testsuite"
	"github.com/pactus-project/pactus/wallet/encrypter"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const tPassword = "super_secret_password"

type testData struct {
	*testsuite.TestSuite

	vault       *Vault
	mnemonic    string
	importedPrv crypto.PrivateKey
}

// setup returns an instances of vault fo testing.
func setup(t *testing.T) *testData {
	t.Helper()

	ts := testsuite.NewTestSuite(t)

	mnemonic, _ := GenerateMnemonic(128)
	_, importedPrv := ts.RandBLSKeyPair()
	vault, err := CreateVaultFromMnemonic(mnemonic, 21888)
	assert.NoError(t, err)

	key, _ := hdkeychain.NewKeyFromString(vault.Purposes.PurposeBLS.XPubAccount)
	assert.False(t, key.IsPrivate())

	// Create some test address
	_, err = vault.NewBLSAccountAddress("addr-1")
	assert.NoError(t, err)
	_, err = vault.NewBLSAccountAddress("addr-2")
	assert.NoError(t, err)
	_, err = vault.NewValidatorAddress("addr-3")
	assert.NoError(t, err)

	assert.NoError(t, vault.ImportPrivateKey("", importedPrv))
	assert.False(t, vault.IsEncrypted())

	opts := []encrypter.Option{
		encrypter.OptionIteration(1),
		encrypter.OptionMemory(1),
		encrypter.OptionParallelism(1),
	}

	err = vault.UpdatePassword("", tPassword, opts...)
	assert.NoError(t, err)
	assert.True(t, vault.IsEncrypted())
	return &testData{
		TestSuite:   ts,
		vault:       vault,
		mnemonic:    mnemonic,
		importedPrv: importedPrv,
	}
}

func TestAddressInfo(t *testing.T) {
	td := setup(t)

	assert.Equal(t, td.vault.AddressCount(), 5)
	infos := td.vault.AddressInfos()
	accountIndex := 0
	validatorIndex := 0
	for _, i := range infos {
		info := td.vault.AddressInfo(i.Address)
		assert.Equal(t, i.Address, info.Address)
		// TODO test me later
		// assert.Equal(t, i.Address, info.PublicKey)

		addr, _ := crypto.AddressFromString(info.Address)
		if info.Path == "" {
			continue
		}

		if addr.IsValidatorAddress() {
			assert.Equal(t, info.Path, fmt.Sprintf("m/12381'/21888'/1'/%d", validatorIndex))
			validatorIndex++
		}

		if addr.IsAccountAddress() {
			assert.Equal(t, info.Path, fmt.Sprintf("m/12381'/21888'/2'/%d", accountIndex))
			accountIndex++
		}
	}

	// Neutered
	neutered := td.vault.Neuter()
	assert.Equal(t, neutered.AddressCount(), 3)
	infos = neutered.AddressInfos()
	accountIndex = 0
	validatorIndex = 0
	for _, i := range infos {
		info := td.vault.AddressInfo(i.Address)
		assert.Equal(t, i.Address, info.Address)
		// TODO test me later
		// assert.Equal(t, i.PublicKey, info.PublicKey)

		if info.Path == "" {
			continue
		}

		addr, _ := crypto.AddressFromString(info.Address)
		if addr.IsValidatorAddress() {
			assert.Equal(t, info.Path, fmt.Sprintf("m/12381'/21888'/1'/%d", validatorIndex))
			validatorIndex++
		}

		if addr.IsAccountAddress() {
			assert.Equal(t, info.Path, fmt.Sprintf("m/12381'/21888'/2'/%d", accountIndex))
			accountIndex++
		}
	}
}

func TestNewBLSAccountAddress(t *testing.T) {
	td := setup(t)

	t.Run("Ok", func(t *testing.T) {
		addr, err := td.vault.NewBLSAccountAddress("new-addr")
		assert.NoError(t, err)
		assert.True(t, td.vault.Contains(addr))
		assert.Equal(t, td.vault.Label(addr), "new-addr")
	})
}

func TestRecover(t *testing.T) {
	td := setup(t)

	t.Run("Invalid mnemonic", func(t *testing.T) {
		_, err := CreateVaultFromMnemonic("invalid mnemonic phrase seed", 21888)
		assert.Error(t, err)
	})

	t.Run("Ok", func(t *testing.T) {
		recovered, err := CreateVaultFromMnemonic(td.mnemonic, 21888)
		assert.NoError(t, err)

		// Recover addresses
		_, err = recovered.NewBLSAccountAddress("addr-1")
		assert.NoError(t, err)
		_, err = recovered.NewBLSAccountAddress("addr-2")
		assert.NoError(t, err)
		_, err = recovered.NewValidatorAddress("addr-3")
		assert.NoError(t, err)

		assert.Equal(t, recovered.Purposes, td.vault.Purposes)
	})
}

func TestGetPrivateKeys(t *testing.T) {
	td := setup(t)

	t.Run("Unknown address", func(t *testing.T) {
		addr := td.RandAccAddress()
		_, err := td.vault.PrivateKeys(tPassword, []string{addr.String()})
		assert.ErrorIs(t, err, NewErrAddressNotFound(addr.String()))
	})

	t.Run("No password", func(t *testing.T) {
		addr := td.vault.AddressInfos()[0].Address
		_, err := td.vault.PrivateKeys("", []string{addr})
		assert.ErrorIs(t, err, encrypter.ErrInvalidPassword)
	})

	t.Run("Invalid password", func(t *testing.T) {
		addr := td.vault.AddressInfos()[0].Address
		_, err := td.vault.PrivateKeys("wrong_password", []string{addr})
		assert.ErrorIs(t, err, encrypter.ErrInvalidPassword)
	})

	t.Run("Check all the private keys", func(t *testing.T) {
		for _, info := range td.vault.AddressInfos() {
			prv, err := td.vault.PrivateKeys(tPassword, []string{info.Address})
			assert.NoError(t, err)
			i := td.vault.AddressInfo(info.Address)
			pub, _ := bls.PublicKeyFromString(i.PublicKey)
			require.True(t, prv[0].PublicKey().EqualsTo(pub))
		}
	})
}

func TestImportPrivateKey(t *testing.T) {
	td := setup(t)

	t.Run("Reimporting private key", func(t *testing.T) {
		err := td.vault.ImportPrivateKey(tPassword, td.importedPrv.(*bls.PrivateKey))
		assert.ErrorIs(t, err, ErrAddressExists)
	})

	t.Run("Invalid password", func(t *testing.T) {
		_, prv := td.RandBLSKeyPair()
		err := td.vault.ImportPrivateKey("invalid-password", prv)
		assert.ErrorIs(t, err, encrypter.ErrInvalidPassword)
	})

	t.Run("Ok", func(t *testing.T) {
		_, prv := td.RandBLSKeyPair()
		assert.NoError(t, td.vault.ImportPrivateKey(tPassword, prv))
		assert.True(t, td.vault.Contains(prv.PublicKeyNative().AccountAddress().String()))
		assert.True(t, td.vault.Contains(prv.PublicKeyNative().ValidatorAddress().String()))
	})
}

func TestGetMnemonic(t *testing.T) {
	td := setup(t)

	t.Run("Invalid password", func(t *testing.T) {
		_, err := td.vault.Mnemonic("invalid-password")
		assert.ErrorIs(t, err, encrypter.ErrInvalidPassword)
	})

	t.Run("No password", func(t *testing.T) {
		_, err := td.vault.Mnemonic("")
		assert.ErrorIs(t, err, encrypter.ErrInvalidPassword)
	})

	t.Run("Ok", func(t *testing.T) {
		m, err := td.vault.Mnemonic(tPassword)
		assert.NoError(t, err)
		assert.Equal(t, m, td.mnemonic)
	})

	t.Run("Neutered wallet", func(t *testing.T) {
		_, err := td.vault.Neuter().Mnemonic("")
		assert.ErrorIs(t, err, ErrNeutered)
	})
}

func TestUpdatePassword(t *testing.T) {
	td := setup(t)

	infos := make([]*AddressInfo, 0, td.vault.AddressCount())
	for _, info := range td.vault.AddressInfos() {
		info := td.vault.AddressInfo(info.Address)
		infos = append(infos, info)
	}

	newPassword := "new-password"
	t.Run("Change password", func(t *testing.T) {
		opts := []encrypter.Option{
			encrypter.OptionIteration(1),
			encrypter.OptionMemory(1),
			encrypter.OptionParallelism(1),
		}

		err := td.vault.UpdatePassword("", newPassword)
		assert.ErrorIs(t, err, encrypter.ErrInvalidPassword)
		err = td.vault.UpdatePassword("invalid-password", newPassword)
		assert.ErrorIs(t, err, encrypter.ErrInvalidPassword)
		assert.NoError(t, td.vault.UpdatePassword(tPassword, newPassword, opts...))
		assert.True(t, td.vault.IsEncrypted())
		for _, info := range infos {
			assert.Equal(t, info, td.vault.AddressInfo(info.Address))
		}
	})

	t.Run("Set empty password for the vault", func(t *testing.T) {
		err := td.vault.UpdatePassword("invalid-password", newPassword)
		assert.ErrorIs(t, err, encrypter.ErrInvalidPassword)
		assert.NoError(t, td.vault.UpdatePassword(newPassword, ""))
		assert.False(t, td.vault.IsEncrypted())
		for _, info := range infos {
			assert.Equal(t, info, td.vault.AddressInfo(info.Address))
		}
	})
}

func TestSetLabel(t *testing.T) {
	td := setup(t)

	t.Run("Set label for unknown address", func(t *testing.T) {
		invAddr := td.RandAccAddress().String()
		err := td.vault.SetLabel(invAddr, "i have label")
		assert.ErrorIs(t, err, NewErrAddressNotFound(invAddr))
		assert.Equal(t, td.vault.Label(invAddr), "")
	})

	t.Run("Update label", func(t *testing.T) {
		testAddr := td.vault.AddressInfos()[0].Address
		err := td.vault.SetLabel(testAddr, "i have label")
		assert.NoError(t, err)
		assert.Equal(t, td.vault.Label(testAddr), "i have label")
	})

	t.Run("Remove label", func(t *testing.T) {
		testAddr := td.vault.AddressInfos()[0].Address
		err := td.vault.SetLabel(testAddr, "")
		assert.NoError(t, err)
		var ok bool
		l := td.vault.Label(testAddr)
		if strings.TrimSpace(l) != "" {
			ok = true
		}
		assert.Empty(t, td.vault.Label(testAddr))
		assert.False(t, ok)
	})
}

func TestNeuter(t *testing.T) {
	td := setup(t)

	neutered := td.vault.Neuter()
	_, err := neutered.Mnemonic(tPassword)
	assert.ErrorIs(t, err, ErrNeutered)

	_, err = neutered.PrivateKeys(tPassword, []string{
		td.RandAccAddress().String(),
	})
	assert.ErrorIs(t, err, ErrNeutered)

	err = neutered.ImportPrivateKey("any", td.importedPrv.(*bls.PrivateKey))
	assert.ErrorIs(t, err, ErrNeutered)

	err = td.vault.Neuter().UpdatePassword("any", "any")
	assert.ErrorIs(t, err, ErrNeutered)
}
