package hipstershop;

import static io.grpc.MethodDescriptor.generateFullMethodName;
import static io.grpc.stub.ClientCalls.asyncBidiStreamingCall;
import static io.grpc.stub.ClientCalls.asyncClientStreamingCall;
import static io.grpc.stub.ClientCalls.asyncServerStreamingCall;
import static io.grpc.stub.ClientCalls.asyncUnaryCall;
import static io.grpc.stub.ClientCalls.blockingServerStreamingCall;
import static io.grpc.stub.ClientCalls.blockingUnaryCall;
import static io.grpc.stub.ClientCalls.futureUnaryCall;
import static io.grpc.stub.ServerCalls.asyncBidiStreamingCall;
import static io.grpc.stub.ServerCalls.asyncClientStreamingCall;
import static io.grpc.stub.ServerCalls.asyncServerStreamingCall;
import static io.grpc.stub.ServerCalls.asyncUnaryCall;
import static io.grpc.stub.ServerCalls.asyncUnimplementedStreamingCall;
import static io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall;

/**
 */
@javax.annotation.Generated(
    value = "by gRPC proto compiler (version 1.10.1)",
    comments = "Source: demo.proto")
public final class CheckoutServiceGrpc {

  private CheckoutServiceGrpc() {}

  public static final String SERVICE_NAME = "hipstershop.CheckoutService";

  // Static method descriptors that strictly reflect the proto.
  @io.grpc.ExperimentalApi("https://github.com/grpc/grpc-java/issues/1901")
  @java.lang.Deprecated // Use {@link #getPlaceOrderMethod()} instead. 
  public static final io.grpc.MethodDescriptor<hipstershop.Demo.PlaceOrderRequest,
      hipstershop.Demo.PlaceOrderResponse> METHOD_PLACE_ORDER = getPlaceOrderMethodHelper();

  private static volatile io.grpc.MethodDescriptor<hipstershop.Demo.PlaceOrderRequest,
      hipstershop.Demo.PlaceOrderResponse> getPlaceOrderMethod;

  @io.grpc.ExperimentalApi("https://github.com/grpc/grpc-java/issues/1901")
  public static io.grpc.MethodDescriptor<hipstershop.Demo.PlaceOrderRequest,
      hipstershop.Demo.PlaceOrderResponse> getPlaceOrderMethod() {
    return getPlaceOrderMethodHelper();
  }

  private static io.grpc.MethodDescriptor<hipstershop.Demo.PlaceOrderRequest,
      hipstershop.Demo.PlaceOrderResponse> getPlaceOrderMethodHelper() {
    io.grpc.MethodDescriptor<hipstershop.Demo.PlaceOrderRequest, hipstershop.Demo.PlaceOrderResponse> getPlaceOrderMethod;
    if ((getPlaceOrderMethod = CheckoutServiceGrpc.getPlaceOrderMethod) == null) {
      synchronized (CheckoutServiceGrpc.class) {
        if ((getPlaceOrderMethod = CheckoutServiceGrpc.getPlaceOrderMethod) == null) {
          CheckoutServiceGrpc.getPlaceOrderMethod = getPlaceOrderMethod = 
              io.grpc.MethodDescriptor.<hipstershop.Demo.PlaceOrderRequest, hipstershop.Demo.PlaceOrderResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(
                  "hipstershop.CheckoutService", "PlaceOrder"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  hipstershop.Demo.PlaceOrderRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  hipstershop.Demo.PlaceOrderResponse.getDefaultInstance()))
                  .setSchemaDescriptor(new CheckoutServiceMethodDescriptorSupplier("PlaceOrder"))
                  .build();
          }
        }
     }
     return getPlaceOrderMethod;
  }

  /**
   * Creates a new async stub that supports all call types for the service
   */
  public static CheckoutServiceStub newStub(io.grpc.Channel channel) {
    return new CheckoutServiceStub(channel);
  }

  /**
   * Creates a new blocking-style stub that supports unary and streaming output calls on the service
   */
  public static CheckoutServiceBlockingStub newBlockingStub(
      io.grpc.Channel channel) {
    return new CheckoutServiceBlockingStub(channel);
  }

  /**
   * Creates a new ListenableFuture-style stub that supports unary calls on the service
   */
  public static CheckoutServiceFutureStub newFutureStub(
      io.grpc.Channel channel) {
    return new CheckoutServiceFutureStub(channel);
  }

  /**
   */
  public static abstract class CheckoutServiceImplBase implements io.grpc.BindableService {

    /**
     */
    public void placeOrder(hipstershop.Demo.PlaceOrderRequest request,
        io.grpc.stub.StreamObserver<hipstershop.Demo.PlaceOrderResponse> responseObserver) {
      asyncUnimplementedUnaryCall(getPlaceOrderMethodHelper(), responseObserver);
    }

    @java.lang.Override public final io.grpc.ServerServiceDefinition bindService() {
      return io.grpc.ServerServiceDefinition.builder(getServiceDescriptor())
          .addMethod(
            getPlaceOrderMethodHelper(),
            asyncUnaryCall(
              new MethodHandlers<
                hipstershop.Demo.PlaceOrderRequest,
                hipstershop.Demo.PlaceOrderResponse>(
                  this, METHODID_PLACE_ORDER)))
          .build();
    }
  }

  /**
   */
  public static final class CheckoutServiceStub extends io.grpc.stub.AbstractStub<CheckoutServiceStub> {
    private CheckoutServiceStub(io.grpc.Channel channel) {
      super(channel);
    }

    private CheckoutServiceStub(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected CheckoutServiceStub build(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      return new CheckoutServiceStub(channel, callOptions);
    }

    /**
     */
    public void placeOrder(hipstershop.Demo.PlaceOrderRequest request,
        io.grpc.stub.StreamObserver<hipstershop.Demo.PlaceOrderResponse> responseObserver) {
      asyncUnaryCall(
          getChannel().newCall(getPlaceOrderMethodHelper(), getCallOptions()), request, responseObserver);
    }
  }

  /**
   */
  public static final class CheckoutServiceBlockingStub extends io.grpc.stub.AbstractStub<CheckoutServiceBlockingStub> {
    private CheckoutServiceBlockingStub(io.grpc.Channel channel) {
      super(channel);
    }

    private CheckoutServiceBlockingStub(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected CheckoutServiceBlockingStub build(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      return new CheckoutServiceBlockingStub(channel, callOptions);
    }

    /**
     */
    public hipstershop.Demo.PlaceOrderResponse placeOrder(hipstershop.Demo.PlaceOrderRequest request) {
      return blockingUnaryCall(
          getChannel(), getPlaceOrderMethodHelper(), getCallOptions(), request);
    }
  }

  /**
   */
  public static final class CheckoutServiceFutureStub extends io.grpc.stub.AbstractStub<CheckoutServiceFutureStub> {
    private CheckoutServiceFutureStub(io.grpc.Channel channel) {
      super(channel);
    }

    private CheckoutServiceFutureStub(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected CheckoutServiceFutureStub build(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      return new CheckoutServiceFutureStub(channel, callOptions);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<hipstershop.Demo.PlaceOrderResponse> placeOrder(
        hipstershop.Demo.PlaceOrderRequest request) {
      return futureUnaryCall(
          getChannel().newCall(getPlaceOrderMethodHelper(), getCallOptions()), request);
    }
  }

  private static final int METHODID_PLACE_ORDER = 0;

  private static final class MethodHandlers<Req, Resp> implements
      io.grpc.stub.ServerCalls.UnaryMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ServerStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ClientStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.BidiStreamingMethod<Req, Resp> {
    private final CheckoutServiceImplBase serviceImpl;
    private final int methodId;

    MethodHandlers(CheckoutServiceImplBase serviceImpl, int methodId) {
      this.serviceImpl = serviceImpl;
      this.methodId = methodId;
    }

    @java.lang.Override
    @java.lang.SuppressWarnings("unchecked")
    public void invoke(Req request, io.grpc.stub.StreamObserver<Resp> responseObserver) {
      switch (methodId) {
        case METHODID_PLACE_ORDER:
          serviceImpl.placeOrder((hipstershop.Demo.PlaceOrderRequest) request,
              (io.grpc.stub.StreamObserver<hipstershop.Demo.PlaceOrderResponse>) responseObserver);
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

  private static abstract class CheckoutServiceBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoFileDescriptorSupplier, io.grpc.protobuf.ProtoServiceDescriptorSupplier {
    CheckoutServiceBaseDescriptorSupplier() {}

    @java.lang.Override
    public com.google.protobuf.Descriptors.FileDescriptor getFileDescriptor() {
      return hipstershop.Demo.getDescriptor();
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.ServiceDescriptor getServiceDescriptor() {
      return getFileDescriptor().findServiceByName("CheckoutService");
    }
  }

  private static final class CheckoutServiceFileDescriptorSupplier
      extends CheckoutServiceBaseDescriptorSupplier {
    CheckoutServiceFileDescriptorSupplier() {}
  }

  private static final class CheckoutServiceMethodDescriptorSupplier
      extends CheckoutServiceBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoMethodDescriptorSupplier {
    private final String methodName;

    CheckoutServiceMethodDescriptorSupplier(String methodName) {
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
      synchronized (CheckoutServiceGrpc.class) {
        result = serviceDescriptor;
        if (result == null) {
          serviceDescriptor = result = io.grpc.ServiceDescriptor.newBuilder(SERVICE_NAME)
              .setSchemaDescriptor(new CheckoutServiceFileDescriptorSupplier())
              .addMethod(getPlaceOrderMethodHelper())
              .build();
        }
      }
    }
    return result;
  }
}
