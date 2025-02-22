package store

import (
	"bytes"
	"errors"
	"sync"

	"github.com/pactus-project/pactus/crypto"
	"github.com/pactus-project/pactus/crypto/bls"
	"github.com/pactus-project/pactus/crypto/hash"
	"github.com/pactus-project/pactus/types/account"
	"github.com/pactus-project/pactus/types/block"
	"github.com/pactus-project/pactus/types/certificate"
	"github.com/pactus-project/pactus/types/tx"
	"github.com/pactus-project/pactus/types/validator"
	"github.com/pactus-project/pactus/util"
	"github.com/pactus-project/pactus/util/encoding"
	"github.com/pactus-project/pactus/util/logger"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/opt"
)

var (
	ErrNotFound  = errors.New("not found")
	ErrBadOffset = errors.New("offset is out of range")
)

const lastStoreVersion = int32(1)

// TODO: add cache for me

var (
	lastInfoKey       = []byte{0x00}
	blockPrefix       = []byte{0x01}
	txPrefix          = []byte{0x03}
	accountPrefix     = []byte{0x05}
	validatorPrefix   = []byte{0x07}
	blockHeightPrefix = []byte{0x09}
	publicKeyPrefix   = []byte{0x0a}
)

func tryGet(db *leveldb.DB, key []byte) ([]byte, error) {
	data, err := db.Get(key, nil)
	if err != nil {
		// Probably key doesn't exist in database
		logger.Trace("database error", "error", err, "key", key)
		return nil, err
	}
	return data, nil
}

type store struct {
	lk sync.RWMutex

	config         *Config
	db             *leveldb.DB
	batch          *leveldb.Batch
	blockStore     *blockStore
	txStore        *txStore
	accountStore   *accountStore
	validatorStore *validatorStore
}

func NewStore(conf *Config) (Store, error) {
	options := &opt.Options{
		Strict:      opt.DefaultStrict,
		Compression: opt.NoCompression,
	}
	db, err := leveldb.OpenFile(conf.StorePath(), options)
	if err != nil {
		return nil, err
	}
	s := &store{
		config:         conf,
		db:             db,
		batch:          new(leveldb.Batch),
		blockStore:     newBlockStore(db),
		txStore:        newTxStore(db),
		accountStore:   newAccountStore(db),
		validatorStore: newValidatorStore(db),
	}
	return s, nil
}

func (s *store) Close() error {
	s.lk.Lock()
	defer s.lk.Unlock()

	return s.db.Close()
}

func (s *store) SaveBlock(blk *block.Block, cert *certificate.Certificate) {
	s.lk.Lock()
	defer s.lk.Unlock()

	height := cert.Height()
	reg := s.blockStore.saveBlock(s.batch, height, blk)
	for i, trx := range blk.Transactions() {
		s.txStore.saveTx(s.batch, trx.ID(), &reg[i])
	}

	// Save last certificate: [version: 4 bytes]+[certificate: variant]
	w := bytes.NewBuffer(make([]byte, 0, 4+cert.SerializeSize()))
	err := encoding.WriteElements(w, lastStoreVersion)
	if err != nil {
		panic(err)
	}
	err = cert.Encode(w)
	if err != nil {
		panic(err)
	}

	s.batch.Put(lastInfoKey, w.Bytes())
}

func (s *store) Block(height uint32) (*CommittedBlock, error) {
	s.lk.Lock()
	defer s.lk.Unlock()

	data, err := s.blockStore.block(height)
	if err != nil {
		return nil, err
	}

	blockHash, err := hash.FromBytes(data[0:hash.HashSize])
	if err != nil {
		return nil, err
	}

	return &CommittedBlock{
		store:     s,
		BlockHash: blockHash,
		Height:    height,
		Data:      data[hash.HashSize:],
	}, nil
}

func (s *store) BlockHeight(h hash.Hash) uint32 {
	s.lk.Lock()
	defer s.lk.Unlock()

	return s.blockStore.blockHeight(h)
}

func (s *store) BlockHash(height uint32) hash.Hash {
	s.lk.Lock()
	defer s.lk.Unlock()

	data, err := s.blockStore.block(height)
	if err == nil {
		blockHash, _ := hash.FromBytes(data[0:hash.HashSize])
		return blockHash
	}

	return hash.UndefHash
}

func (s *store) PublicKey(addr crypto.Address) (*bls.PublicKey, error) {
	bs, err := tryGet(s.db, publicKeyKey(addr))
	if err != nil {
		return nil, err
	}
	pubKey, err := bls.PublicKeyFromBytes(bs)
	if err != nil {
		return nil, err
	}

	return pubKey, err
}

func (s *store) Transaction(id tx.ID) (*CommittedTx, error) {
	s.lk.Lock()
	defer s.lk.Unlock()

	pos, err := s.txStore.tx(id)
	if err != nil {
		return nil, err
	}
	data, err := s.blockStore.block(pos.height)
	if err != nil {
		return nil, err
	}
	start := pos.offset
	end := pos.offset + pos.length
	if end > uint32(len(data)) {
		return nil, ErrBadOffset
	}
	blockTime := util.SliceToUint32(data[hash.HashSize+1 : hash.HashSize+5])

	return &CommittedTx{
		store:     s,
		TxID:      id,
		Height:    pos.height,
		BlockTime: blockTime,
		Data:      data[start:end],
	}, nil
}

// TODO implement Dequeue for this function, for the better performance.
func (s *store) AnyRecentTransaction(id tx.ID) bool {
	s.lk.Lock()
	defer s.lk.Unlock()

	pos, _ := s.txStore.tx(id)

	return pos != nil
}

func (s *store) HasAccount(addr crypto.Address) bool {
	s.lk.Lock()
	defer s.lk.Unlock()

	return s.accountStore.hasAccount(addr)
}

func (s *store) Account(addr crypto.Address) (*account.Account, error) {
	s.lk.Lock()
	defer s.lk.Unlock()

	return s.accountStore.account(addr)
}

func (s *store) TotalAccounts() int32 {
	s.lk.Lock()
	defer s.lk.Unlock()

	return s.accountStore.total
}

func (s *store) IterateAccounts(consumer func(crypto.Address, *account.Account) (stop bool)) {
	s.lk.Lock()
	defer s.lk.Unlock()

	s.accountStore.iterateAccounts(consumer)
}

func (s *store) UpdateAccount(addr crypto.Address, acc *account.Account) {
	s.lk.Lock()
	defer s.lk.Unlock()

	s.accountStore.updateAccount(s.batch, addr, acc)
}

func (s *store) HasValidator(addr crypto.Address) bool {
	s.lk.Lock()
	defer s.lk.Unlock()

	return s.validatorStore.hasValidator(addr)
}

func (s *store) ValidatorAddresses() []crypto.Address {
	s.lk.RLock()
	defer s.lk.RUnlock()

	return s.validatorStore.ValidatorAddresses()
}

func (s *store) Validator(addr crypto.Address) (*validator.Validator, error) {
	s.lk.Lock()
	defer s.lk.Unlock()

	return s.validatorStore.validator(addr)
}

func (s *store) ValidatorByNumber(num int32) (*validator.Validator, error) {
	s.lk.Lock()
	defer s.lk.Unlock()

	return s.validatorStore.validatorByNumber(num)
}

func (s *store) TotalValidators() int32 {
	s.lk.Lock()
	defer s.lk.Unlock()

	return s.validatorStore.total
}

func (s *store) IterateValidators(consumer func(*validator.Validator) (stop bool)) {
	s.lk.Lock()
	defer s.lk.Unlock()

	s.validatorStore.iterateValidators(consumer)
}

func (s *store) UpdateValidator(acc *validator.Validator) {
	s.lk.Lock()
	defer s.lk.Unlock()

	s.validatorStore.updateValidator(s.batch, acc)
}

func (s *store) LastCertificate() *certificate.Certificate {
	s.lk.Lock()
	defer s.lk.Unlock()

	data, _ := tryGet(s.db, lastInfoKey)
	if data == nil {
		// Genesis block
		return nil
	}
	r := bytes.NewReader(data)
	version := int32(0)
	cert := new(certificate.Certificate)
	err := encoding.ReadElements(r, &version)
	if err != nil {
		return nil
	}
	err = cert.Decode(r)
	if err != nil {
		return nil
	}
	return cert
}

func (s *store) WriteBatch() error {
	s.lk.Lock()
	defer s.lk.Unlock()

	if err := s.db.Write(s.batch, nil); err != nil {
		// TODO: Should we panic here?
		// The store is unreliable if the stored data does not match the cached data.
		return err
	}
	s.batch.Reset()
	return nil
}
