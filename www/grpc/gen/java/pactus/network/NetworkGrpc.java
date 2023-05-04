package pactus.network;

import static io.grpc.MethodDescriptor.generateFullMethodName;

/**
 */
@javax.annotation.Generated(
    value = "by gRPC proto compiler (version 1.50.2)",
    comments = "Source: network.proto")
@io.grpc.stub.annotations.GrpcGenerated
public final class NetworkGrpc {

  private NetworkGrpc() {}

  public static final String SERVICE_NAME = "pactus.Network";

  // Static method descriptors that strictly reflect the proto.
  private static volatile io.grpc.MethodDescriptor<pactus.network.NetworkOuterClass.GetNetworkInfoRequest,
      pactus.network.NetworkOuterClass.GetNetworkInfoResponse> getGetNetworkInfoMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetNetworkInfo",
      requestType = pactus.network.NetworkOuterClass.GetNetworkInfoRequest.class,
      responseType = pactus.network.NetworkOuterClass.GetNetworkInfoResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<pactus.network.NetworkOuterClass.GetNetworkInfoRequest,
      pactus.network.NetworkOuterClass.GetNetworkInfoResponse> getGetNetworkInfoMethod() {
    io.grpc.MethodDescriptor<pactus.network.NetworkOuterClass.GetNetworkInfoRequest, pactus.network.NetworkOuterClass.GetNetworkInfoResponse> getGetNetworkInfoMethod;
    if ((getGetNetworkInfoMethod = NetworkGrpc.getGetNetworkInfoMethod) == null) {
      synchronized (NetworkGrpc.class) {
        if ((getGetNetworkInfoMethod = NetworkGrpc.getGetNetworkInfoMethod) == null) {
          NetworkGrpc.getGetNetworkInfoMethod = getGetNetworkInfoMethod =
              io.grpc.MethodDescriptor.<pactus.network.NetworkOuterClass.GetNetworkInfoRequest, pactus.network.NetworkOuterClass.GetNetworkInfoResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetNetworkInfo"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  pactus.network.NetworkOuterClass.GetNetworkInfoRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  pactus.network.NetworkOuterClass.GetNetworkInfoResponse.getDefaultInstance()))
              .setSchemaDescriptor(new NetworkMethodDescriptorSupplier("GetNetworkInfo"))
              .build();
        }
      }
    }
    return getGetNetworkInfoMethod;
  }

  private static volatile io.grpc.MethodDescriptor<pactus.network.NetworkOuterClass.GetPeerInfoRequest,
      pactus.network.NetworkOuterClass.GetPeerInfoResponse> getGetPeerInfoMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetPeerInfo",
      requestType = pactus.network.NetworkOuterClass.GetPeerInfoRequest.class,
      responseType = pactus.network.NetworkOuterClass.GetPeerInfoResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<pactus.network.NetworkOuterClass.GetPeerInfoRequest,
      pactus.network.NetworkOuterClass.GetPeerInfoResponse> getGetPeerInfoMethod() {
    io.grpc.MethodDescriptor<pactus.network.NetworkOuterClass.GetPeerInfoRequest, pactus.network.NetworkOuterClass.GetPeerInfoResponse> getGetPeerInfoMethod;
    if ((getGetPeerInfoMethod = NetworkGrpc.getGetPeerInfoMethod) == null) {
      synchronized (NetworkGrpc.class) {
        if ((getGetPeerInfoMethod = NetworkGrpc.getGetPeerInfoMethod) == null) {
          NetworkGrpc.getGetPeerInfoMethod = getGetPeerInfoMethod =
              io.grpc.MethodDescriptor.<pactus.network.NetworkOuterClass.GetPeerInfoRequest, pactus.network.NetworkOuterClass.GetPeerInfoResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetPeerInfo"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  pactus.network.NetworkOuterClass.GetPeerInfoRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  pactus.network.NetworkOuterClass.GetPeerInfoResponse.getDefaultInstance()))
              .setSchemaDescriptor(new NetworkMethodDescriptorSupplier("GetPeerInfo"))
              .build();
        }
      }
    }
    return getGetPeerInfoMethod;
  }

  /**
   * Creates a new async stub that supports all call types for the service
   */
  public static NetworkStub newStub(io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<NetworkStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<NetworkStub>() {
        @java.lang.Override
        public NetworkStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new NetworkStub(channel, callOptions);
        }
      };
    return NetworkStub.newStub(factory, channel);
  }

  /**
   * Creates a new blocking-style stub that supports unary and streaming output calls on the service
   */
  public static NetworkBlockingStub newBlockingStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<NetworkBlockingStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<NetworkBlockingStub>() {
        @java.lang.Override
        public NetworkBlockingStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new NetworkBlockingStub(channel, callOptions);
        }
      };
    return NetworkBlockingStub.newStub(factory, channel);
  }

  /**
   * Creates a new ListenableFuture-style stub that supports unary calls on the service
   */
  public static NetworkFutureStub newFutureStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<NetworkFutureStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<NetworkFutureStub>() {
        @java.lang.Override
        public NetworkFutureStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new NetworkFutureStub(channel, callOptions);
        }
      };
    return NetworkFutureStub.newStub(factory, channel);
  }

  /**
   */
  public static abstract class NetworkImplBase implements io.grpc.BindableService {

    /**
     */
    public void getNetworkInfo(pactus.network.NetworkOuterClass.GetNetworkInfoRequest request,
        io.grpc.stub.StreamObserver<pactus.network.NetworkOuterClass.GetNetworkInfoResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetNetworkInfoMethod(), responseObserver);
    }

    /**
     */
    public void getPeerInfo(pactus.network.NetworkOuterClass.GetPeerInfoRequest request,
        io.grpc.stub.StreamObserver<pactus.network.NetworkOuterClass.GetPeerInfoResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetPeerInfoMethod(), responseObserver);
    }

    @java.lang.Override public final io.grpc.ServerServiceDefinition bindService() {
      return io.grpc.ServerServiceDefinition.builder(getServiceDescriptor())
          .addMethod(
            getGetNetworkInfoMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                pactus.network.NetworkOuterClass.GetNetworkInfoRequest,
                pactus.network.NetworkOuterClass.GetNetworkInfoResponse>(
                  this, METHODID_GET_NETWORK_INFO)))
          .addMethod(
            getGetPeerInfoMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                pactus.network.NetworkOuterClass.GetPeerInfoRequest,
                pactus.network.NetworkOuterClass.GetPeerInfoResponse>(
                  this, METHODID_GET_PEER_INFO)))
          .build();
    }
  }

  /**
   */
  public static final class NetworkStub extends io.grpc.stub.AbstractAsyncStub<NetworkStub> {
    private NetworkStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected NetworkStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new NetworkStub(channel, callOptions);
    }

    /**
     */
    public void getNetworkInfo(pactus.network.NetworkOuterClass.GetNetworkInfoRequest request,
        io.grpc.stub.StreamObserver<pactus.network.NetworkOuterClass.GetNetworkInfoResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetNetworkInfoMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void getPeerInfo(pactus.network.NetworkOuterClass.GetPeerInfoRequest request,
        io.grpc.stub.StreamObserver<pactus.network.NetworkOuterClass.GetPeerInfoResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetPeerInfoMethod(), getCallOptions()), request, responseObserver);
    }
  }

  /**
   */
  public static final class NetworkBlockingStub extends io.grpc.stub.AbstractBlockingStub<NetworkBlockingStub> {
    private NetworkBlockingStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected NetworkBlockingStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new NetworkBlockingStub(channel, callOptions);
    }

    /**
     */
    public pactus.network.NetworkOuterClass.GetNetworkInfoResponse getNetworkInfo(pactus.network.NetworkOuterClass.GetNetworkInfoRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetNetworkInfoMethod(), getCallOptions(), request);
    }

    /**
     */
    public pactus.network.NetworkOuterClass.GetPeerInfoResponse getPeerInfo(pactus.network.NetworkOuterClass.GetPeerInfoRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetPeerInfoMethod(), getCallOptions(), request);
    }
  }

  /**
   */
  public static final class NetworkFutureStub extends io.grpc.stub.AbstractFutureStub<NetworkFutureStub> {
    private NetworkFutureStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected NetworkFutureStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new NetworkFutureStub(channel, callOptions);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<pactus.network.NetworkOuterClass.GetNetworkInfoResponse> getNetworkInfo(
        pactus.network.NetworkOuterClass.GetNetworkInfoRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetNetworkInfoMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<pactus.network.NetworkOuterClass.GetPeerInfoResponse> getPeerInfo(
        pactus.network.NetworkOuterClass.GetPeerInfoRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetPeerInfoMethod(), getCallOptions()), request);
    }
  }

  private static final int METHODID_GET_NETWORK_INFO = 0;
  private static final int METHODID_GET_PEER_INFO = 1;

  private static final class MethodHandlers<Req, Resp> implements
      io.grpc.stub.ServerCalls.UnaryMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ServerStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ClientStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.BidiStreamingMethod<Req, Resp> {
    private final NetworkImplBase serviceImpl;
    private final int methodId;

    MethodHandlers(NetworkImplBase serviceImpl, int methodId) {
      this.serviceImpl = serviceImpl;
      this.methodId = methodId;
    }

    @java.lang.Override
    @java.lang.SuppressWarnings("unchecked")
    public void invoke(Req request, io.grpc.stub.StreamObserver<Resp> responseObserver) {
      switch (methodId) {
        case METHODID_GET_NETWORK_INFO:
          serviceImpl.getNetworkInfo((pactus.network.NetworkOuterClass.GetNetworkInfoRequest) request,
              (io.grpc.stub.StreamObserver<pactus.network.NetworkOuterClass.GetNetworkInfoResponse>) responseObserver);
          break;
        case METHODID_GET_PEER_INFO:
          serviceImpl.getPeerInfo((pactus.network.NetworkOuterClass.GetPeerInfoRequest) request,
              (io.grpc.stub.StreamObserver<pactus.network.NetworkOuterClass.GetPeerInfoResponse>) responseObserver);
          break;
        default:
          throw new AssertionError();
      }
    }

    @java.lang.Override
    @java.lang.SuppressWarnings("unchecked")
    public io.grpc.stub.StreamObserver<Req> invoke(
        io.grpc.stub.StreamObserver<Resp> responseObserver) {
      switch (methodId) {
        default:
          throw new AssertionError();
      }
    }
  }

  private static abstract class NetworkBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoFileDescriptorSupplier, io.grpc.protobuf.ProtoServiceDescriptorSupplier {
    NetworkBaseDescriptorSupplier() {}

    @java.lang.Override
    public com.google.protobuf.Descriptors.FileDescriptor getFileDescriptor() {
      return pactus.network.NetworkOuterClass.getDescriptor();
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.ServiceDescriptor getServiceDescriptor() {
      return getFileDescriptor().findServiceByName("Network");
    }
  }

  private static final class NetworkFileDescriptorSupplier
      extends NetworkBaseDescriptorSupplier {
    NetworkFileDescriptorSupplier() {}
  }

  private static final class NetworkMethodDescriptorSupplier
      extends NetworkBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoMethodDescriptorSupplier {
    private final String methodName;

    NetworkMethodDescriptorSupplier(String methodName) {
      this.methodName = methodName;
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.MethodDescriptor getMethodDescriptor() {
      return getServiceDescriptor().findMethodByName(methodName);
    }
  }

  private static volatile io.grpc.ServiceDescriptor serviceDescriptor;

  public static io.grpc.ServiceDescriptor getServiceDescriptor() {
    io.grpc.ServiceDescriptor result = serviceDescriptor;
    if (result == null) {
      synchronized (NetworkGrpc.class) {
        result = serviceDescriptor;
        if (result == null) {
          serviceDescriptor = result = io.grpc.ServiceDescriptor.newBuilder(SERVICE_NAME)
              .setSchemaDescriptor(new NetworkFileDescriptorSupplier())
              .addMethod(getGetNetworkInfoMethod())
              .addMethod(getGetPeerInfoMethod())
              .build();
        }
      }
    }
    return result;
  }
}