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
public final class ShippingServiceGrpc {

  private ShippingServiceGrpc() {}

  public static final String SERVICE_NAME = "hipstershop.ShippingService";

  // Static method descriptors that strictly reflect the proto.
  @io.grpc.ExperimentalApi("https://github.com/grpc/grpc-java/issues/1901")
  @java.lang.Deprecated // Use {@link #getGetQuoteMethod()} instead. 
  public static final io.grpc.MethodDescriptor<hipstershop.Demo.GetQuoteRequest,
      hipstershop.Demo.GetQuoteResponse> METHOD_GET_QUOTE = getGetQuoteMethodHelper();

  private static volatile io.grpc.MethodDescriptor<hipstershop.Demo.GetQuoteRequest,
      hipstershop.Demo.GetQuoteResponse> getGetQuoteMethod;

  @io.grpc.ExperimentalApi("https://github.com/grpc/grpc-java/issues/1901")
  public static io.grpc.MethodDescriptor<hipstershop.Demo.GetQuoteRequest,
      hipstershop.Demo.GetQuoteResponse> getGetQuoteMethod() {
    return getGetQuoteMethodHelper();
  }

  private static io.grpc.MethodDescriptor<hipstershop.Demo.GetQuoteRequest,
      hipstershop.Demo.GetQuoteResponse> getGetQuoteMethodHelper() {
    io.grpc.MethodDescriptor<hipstershop.Demo.GetQuoteRequest, hipstershop.Demo.GetQuoteResponse> getGetQuoteMethod;
    if ((getGetQuoteMethod = ShippingServiceGrpc.getGetQuoteMethod) == null) {
      synchronized (ShippingServiceGrpc.class) {
        if ((getGetQuoteMethod = ShippingServiceGrpc.getGetQuoteMethod) == null) {
          ShippingServiceGrpc.getGetQuoteMethod = getGetQuoteMethod = 
              io.grpc.MethodDescriptor.<hipstershop.Demo.GetQuoteRequest, hipstershop.Demo.GetQuoteResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(
                  "hipstershop.ShippingService", "GetQuote"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  hipstershop.Demo.GetQuoteRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  hipstershop.Demo.GetQuoteResponse.getDefaultInstance()))
                  .setSchemaDescriptor(new ShippingServiceMethodDescriptorSupplier("GetQuote"))
                  .build();
          }
        }
     }
     return getGetQuoteMethod;
  }
  @io.grpc.ExperimentalApi("https://github.com/grpc/grpc-java/issues/1901")
  @java.lang.Deprecated // Use {@link #getShipOrderMethod()} instead. 
  public static final io.grpc.MethodDescriptor<hipstershop.Demo.ShipOrderRequest,
      hipstershop.Demo.ShipOrderResponse> METHOD_SHIP_ORDER = getShipOrderMethodHelper();

  private static volatile io.grpc.MethodDescriptor<hipstershop.Demo.ShipOrderRequest,
      hipstershop.Demo.ShipOrderResponse> getShipOrderMethod;

  @io.grpc.ExperimentalApi("https://github.com/grpc/grpc-java/issues/1901")
  public static io.grpc.MethodDescriptor<hipstershop.Demo.ShipOrderRequest,
      hipstershop.Demo.ShipOrderResponse> getShipOrderMethod() {
    return getShipOrderMethodHelper();
  }

  private static io.grpc.MethodDescriptor<hipstershop.Demo.ShipOrderRequest,
      hipstershop.Demo.ShipOrderResponse> getShipOrderMethodHelper() {
    io.grpc.MethodDescriptor<hipstershop.Demo.ShipOrderRequest, hipstershop.Demo.ShipOrderResponse> getShipOrderMethod;
    if ((getShipOrderMethod = ShippingServiceGrpc.getShipOrderMethod) == null) {
      synchronized (ShippingServiceGrpc.class) {
        if ((getShipOrderMethod = ShippingServiceGrpc.getShipOrderMethod) == null) {
          ShippingServiceGrpc.getShipOrderMethod = getShipOrderMethod = 
              io.grpc.MethodDescriptor.<hipstershop.Demo.ShipOrderRequest, hipstershop.Demo.ShipOrderResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(
                  "hipstershop.ShippingService", "ShipOrder"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  hipstershop.Demo.ShipOrderRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  hipstershop.Demo.ShipOrderResponse.getDefaultInstance()))
                  .setSchemaDescriptor(new ShippingServiceMethodDescriptorSupplier("ShipOrder"))
                  .build();
          }
        }
     }
     return getShipOrderMethod;
  }

  /**
   * Creates a new async stub that supports all call types for the service
   */
  public static ShippingServiceStub newStub(io.grpc.Channel channel) {
    return new ShippingServiceStub(channel);
  }

  /**
   * Creates a new blocking-style stub that supports unary and streaming output calls on the service
   */
  public static ShippingServiceBlockingStub newBlockingStub(
      io.grpc.Channel channel) {
    return new ShippingServiceBlockingStub(channel);
  }

  /**
   * Creates a new ListenableFuture-style stub that supports unary calls on the service
   */
  public static ShippingServiceFutureStub newFutureStub(
      io.grpc.Channel channel) {
    return new ShippingServiceFutureStub(channel);
  }

  /**
   */
  public static abstract class ShippingServiceImplBase implements io.grpc.BindableService {

    /**
     */
    public void getQuote(hipstershop.Demo.GetQuoteRequest request,
        io.grpc.stub.StreamObserver<hipstershop.Demo.GetQuoteResponse> responseObserver) {
      asyncUnimplementedUnaryCall(getGetQuoteMethodHelper(), responseObserver);
    }

    /**
     */
    public void shipOrder(hipstershop.Demo.ShipOrderRequest request,
        io.grpc.stub.StreamObserver<hipstershop.Demo.ShipOrderResponse> responseObserver) {
      asyncUnimplementedUnaryCall(getShipOrderMethodHelper(), responseObserver);
    }

    @java.lang.Override public final io.grpc.ServerServiceDefinition bindService() {
      return io.grpc.ServerServiceDefinition.builder(getServiceDescriptor())
          .addMethod(
            getGetQuoteMethodHelper(),
            asyncUnaryCall(
              new MethodHandlers<
                hipstershop.Demo.GetQuoteRequest,
                hipstershop.Demo.GetQuoteResponse>(
                  this, METHODID_GET_QUOTE)))
          .addMethod(
            getShipOrderMethodHelper(),
            asyncUnaryCall(
              new MethodHandlers<
                hipstershop.Demo.ShipOrderRequest,
                hipstershop.Demo.ShipOrderResponse>(
                  this, METHODID_SHIP_ORDER)))
          .build();
    }
  }

  /**
   */
  public static final class ShippingServiceStub extends io.grpc.stub.AbstractStub<ShippingServiceStub> {
    private ShippingServiceStub(io.grpc.Channel channel) {
      super(channel);
    }

    private ShippingServiceStub(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected ShippingServiceStub build(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      return new ShippingServiceStub(channel, callOptions);
    }

    /**
     */
    public void getQuote(hipstershop.Demo.GetQuoteRequest request,
        io.grpc.stub.StreamObserver<hipstershop.Demo.GetQuoteResponse> responseObserver) {
      asyncUnaryCall(
          getChannel().newCall(getGetQuoteMethodHelper(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void shipOrder(hipstershop.Demo.ShipOrderRequest request,
        io.grpc.stub.StreamObserver<hipstershop.Demo.ShipOrderResponse> responseObserver) {
      asyncUnaryCall(
          getChannel().newCall(getShipOrderMethodHelper(), getCallOptions()), request, responseObserver);
    }
  }

  /**
   */
  public static final class ShippingServiceBlockingStub extends io.grpc.stub.AbstractStub<ShippingServiceBlockingStub> {
    private ShippingServiceBlockingStub(io.grpc.Channel channel) {
      super(channel);
    }

    private ShippingServiceBlockingStub(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected ShippingServiceBlockingStub build(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      return new ShippingServiceBlockingStub(channel, callOptions);
    }

    /**
     */
    public hipstershop.Demo.GetQuoteResponse getQuote(hipstershop.Demo.GetQuoteRequest request) {
      return blockingUnaryCall(
          getChannel(), getGetQuoteMethodHelper(), getCallOptions(), request);
    }

    /**
     */
    public hipstershop.Demo.ShipOrderResponse shipOrder(hipstershop.Demo.ShipOrderRequest request) {
      return blockingUnaryCall(
          getChannel(), getShipOrderMethodHelper(), getCallOptions(), request);
    }
  }

  /**
   */
  public static final class ShippingServiceFutureStub extends io.grpc.stub.AbstractStub<ShippingServiceFutureStub> {
    private ShippingServiceFutureStub(io.grpc.Channel channel) {
      super(channel);
    }

    private ShippingServiceFutureStub(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected ShippingServiceFutureStub build(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      return new ShippingServiceFutureStub(channel, callOptions);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<hipstershop.Demo.GetQuoteResponse> getQuote(
        hipstershop.Demo.GetQuoteRequest request) {
      return futureUnaryCall(
          getChannel().newCall(getGetQuoteMethodHelper(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<hipstershop.Demo.ShipOrderResponse> shipOrder(
        hipstershop.Demo.ShipOrderRequest request) {
      return futureUnaryCall(
          getChannel().newCall(getShipOrderMethodHelper(), getCallOptions()), request);
    }
  }

  private static final int METHODID_GET_QUOTE = 0;
  private static final int METHODID_SHIP_ORDER = 1;

  private static final class MethodHandlers<Req, Resp> implements
      io.grpc.stub.ServerCalls.UnaryMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ServerStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ClientStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.BidiStreamingMethod<Req, Resp> {
    private final ShippingServiceImplBase serviceImpl;
    private final int methodId;

    MethodHandlers(ShippingServiceImplBase serviceImpl, int methodId) {
      this.serviceImpl = serviceImpl;
      this.methodId = methodId;
    }

    @java.lang.Override
    @java.lang.SuppressWarnings("unchecked")
    public void invoke(Req request, io.grpc.stub.StreamObserver<Resp> responseObserver) {
      switch (methodId) {
        case METHODID_GET_QUOTE:
          serviceImpl.getQuote((hipstershop.Demo.GetQuoteRequest) request,
              (io.grpc.stub.StreamObserver<hipstershop.Demo.GetQuoteResponse>) responseObserver);
          break;
        case METHODID_SHIP_ORDER:
          serviceImpl.shipOrder((hipstershop.Demo.ShipOrderRequest) request,
              (io.grpc.stub.StreamObserver<hipstershop.Demo.ShipOrderResponse>) responseObserver);
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

  private static abstract class ShippingServiceBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoFileDescriptorSupplier, io.grpc.protobuf.ProtoServiceDescriptorSupplier {
    ShippingServiceBaseDescriptorSupplier() {}

    @java.lang.Override
    public com.google.protobuf.Descriptors.FileDescriptor getFileDescriptor() {
      return hipstershop.Demo.getDescriptor();
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.ServiceDescriptor getServiceDescriptor() {
      return getFileDescriptor().findServiceByName("ShippingService");
    }
  }

  private static final class ShippingServiceFileDescriptorSupplier
      extends ShippingServiceBaseDescriptorSupplier {
    ShippingServiceFileDescriptorSupplier() {}
  }

  private static final class ShippingServiceMethodDescriptorSupplier
      extends ShippingServiceBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoMethodDescriptorSupplier {
    private final String methodName;

    ShippingServiceMethodDescriptorSupplier(String methodName) {
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
      synchronized (ShippingServiceGrpc.class) {
        result = serviceDescriptor;
        if (result == null) {
          serviceDescriptor = result = io.grpc.ServiceDescriptor.newBuilder(SERVICE_NAME)
              .setSchemaDescriptor(new ShippingServiceFileDescriptorSupplier())
              .addMethod(getGetQuoteMethodHelper())
              .addMethod(getShipOrderMethodHelper())
              .build();
        }
      }
    }
    return result;
  }
}
