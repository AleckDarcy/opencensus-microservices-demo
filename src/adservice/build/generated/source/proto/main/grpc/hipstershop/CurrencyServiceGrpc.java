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
public final class CurrencyServiceGrpc {

  private CurrencyServiceGrpc() {}

  public static final String SERVICE_NAME = "hipstershop.CurrencyService";

  // Static method descriptors that strictly reflect the proto.
  @io.grpc.ExperimentalApi("https://github.com/grpc/grpc-java/issues/1901")
  @java.lang.Deprecated // Use {@link #getGetSupportedCurrenciesMethod()} instead. 
  public static final io.grpc.MethodDescriptor<hipstershop.Demo.GetSupportedCurrenciesRequest,
      hipstershop.Demo.GetSupportedCurrenciesResponse> METHOD_GET_SUPPORTED_CURRENCIES = getGetSupportedCurrenciesMethodHelper();

  private static volatile io.grpc.MethodDescriptor<hipstershop.Demo.GetSupportedCurrenciesRequest,
      hipstershop.Demo.GetSupportedCurrenciesResponse> getGetSupportedCurrenciesMethod;

  @io.grpc.ExperimentalApi("https://github.com/grpc/grpc-java/issues/1901")
  public static io.grpc.MethodDescriptor<hipstershop.Demo.GetSupportedCurrenciesRequest,
      hipstershop.Demo.GetSupportedCurrenciesResponse> getGetSupportedCurrenciesMethod() {
    return getGetSupportedCurrenciesMethodHelper();
  }

  private static io.grpc.MethodDescriptor<hipstershop.Demo.GetSupportedCurrenciesRequest,
      hipstershop.Demo.GetSupportedCurrenciesResponse> getGetSupportedCurrenciesMethodHelper() {
    io.grpc.MethodDescriptor<hipstershop.Demo.GetSupportedCurrenciesRequest, hipstershop.Demo.GetSupportedCurrenciesResponse> getGetSupportedCurrenciesMethod;
    if ((getGetSupportedCurrenciesMethod = CurrencyServiceGrpc.getGetSupportedCurrenciesMethod) == null) {
      synchronized (CurrencyServiceGrpc.class) {
        if ((getGetSupportedCurrenciesMethod = CurrencyServiceGrpc.getGetSupportedCurrenciesMethod) == null) {
          CurrencyServiceGrpc.getGetSupportedCurrenciesMethod = getGetSupportedCurrenciesMethod = 
              io.grpc.MethodDescriptor.<hipstershop.Demo.GetSupportedCurrenciesRequest, hipstershop.Demo.GetSupportedCurrenciesResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(
                  "hipstershop.CurrencyService", "GetSupportedCurrencies"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  hipstershop.Demo.GetSupportedCurrenciesRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  hipstershop.Demo.GetSupportedCurrenciesResponse.getDefaultInstance()))
                  .setSchemaDescriptor(new CurrencyServiceMethodDescriptorSupplier("GetSupportedCurrencies"))
                  .build();
          }
        }
     }
     return getGetSupportedCurrenciesMethod;
  }
  @io.grpc.ExperimentalApi("https://github.com/grpc/grpc-java/issues/1901")
  @java.lang.Deprecated // Use {@link #getConvertMethod()} instead. 
  public static final io.grpc.MethodDescriptor<hipstershop.Demo.CurrencyConversionRequest,
      hipstershop.Demo.Money> METHOD_CONVERT = getConvertMethodHelper();

  private static volatile io.grpc.MethodDescriptor<hipstershop.Demo.CurrencyConversionRequest,
      hipstershop.Demo.Money> getConvertMethod;

  @io.grpc.ExperimentalApi("https://github.com/grpc/grpc-java/issues/1901")
  public static io.grpc.MethodDescriptor<hipstershop.Demo.CurrencyConversionRequest,
      hipstershop.Demo.Money> getConvertMethod() {
    return getConvertMethodHelper();
  }

  private static io.grpc.MethodDescriptor<hipstershop.Demo.CurrencyConversionRequest,
      hipstershop.Demo.Money> getConvertMethodHelper() {
    io.grpc.MethodDescriptor<hipstershop.Demo.CurrencyConversionRequest, hipstershop.Demo.Money> getConvertMethod;
    if ((getConvertMethod = CurrencyServiceGrpc.getConvertMethod) == null) {
      synchronized (CurrencyServiceGrpc.class) {
        if ((getConvertMethod = CurrencyServiceGrpc.getConvertMethod) == null) {
          CurrencyServiceGrpc.getConvertMethod = getConvertMethod = 
              io.grpc.MethodDescriptor.<hipstershop.Demo.CurrencyConversionRequest, hipstershop.Demo.Money>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(
                  "hipstershop.CurrencyService", "Convert"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  hipstershop.Demo.CurrencyConversionRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  hipstershop.Demo.Money.getDefaultInstance()))
                  .setSchemaDescriptor(new CurrencyServiceMethodDescriptorSupplier("Convert"))
                  .build();
          }
        }
     }
     return getConvertMethod;
  }

  /**
   * Creates a new async stub that supports all call types for the service
   */
  public static CurrencyServiceStub newStub(io.grpc.Channel channel) {
    return new CurrencyServiceStub(channel);
  }

  /**
   * Creates a new blocking-style stub that supports unary and streaming output calls on the service
   */
  public static CurrencyServiceBlockingStub newBlockingStub(
      io.grpc.Channel channel) {
    return new CurrencyServiceBlockingStub(channel);
  }

  /**
   * Creates a new ListenableFuture-style stub that supports unary calls on the service
   */
  public static CurrencyServiceFutureStub newFutureStub(
      io.grpc.Channel channel) {
    return new CurrencyServiceFutureStub(channel);
  }

  /**
   */
  public static abstract class CurrencyServiceImplBase implements io.grpc.BindableService {

    /**
     */
    public void getSupportedCurrencies(hipstershop.Demo.GetSupportedCurrenciesRequest request,
        io.grpc.stub.StreamObserver<hipstershop.Demo.GetSupportedCurrenciesResponse> responseObserver) {
      asyncUnimplementedUnaryCall(getGetSupportedCurrenciesMethodHelper(), responseObserver);
    }

    /**
     */
    public void convert(hipstershop.Demo.CurrencyConversionRequest request,
        io.grpc.stub.StreamObserver<hipstershop.Demo.Money> responseObserver) {
      asyncUnimplementedUnaryCall(getConvertMethodHelper(), responseObserver);
    }

    @java.lang.Override public final io.grpc.ServerServiceDefinition bindService() {
      return io.grpc.ServerServiceDefinition.builder(getServiceDescriptor())
          .addMethod(
            getGetSupportedCurrenciesMethodHelper(),
            asyncUnaryCall(
              new MethodHandlers<
                hipstershop.Demo.GetSupportedCurrenciesRequest,
                hipstershop.Demo.GetSupportedCurrenciesResponse>(
                  this, METHODID_GET_SUPPORTED_CURRENCIES)))
          .addMethod(
            getConvertMethodHelper(),
            asyncUnaryCall(
              new MethodHandlers<
                hipstershop.Demo.CurrencyConversionRequest,
                hipstershop.Demo.Money>(
                  this, METHODID_CONVERT)))
          .build();
    }
  }

  /**
   */
  public static final class CurrencyServiceStub extends io.grpc.stub.AbstractStub<CurrencyServiceStub> {
    private CurrencyServiceStub(io.grpc.Channel channel) {
      super(channel);
    }

    private CurrencyServiceStub(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected CurrencyServiceStub build(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      return new CurrencyServiceStub(channel, callOptions);
    }

    /**
     */
    public void getSupportedCurrencies(hipstershop.Demo.GetSupportedCurrenciesRequest request,
        io.grpc.stub.StreamObserver<hipstershop.Demo.GetSupportedCurrenciesResponse> responseObserver) {
      asyncUnaryCall(
          getChannel().newCall(getGetSupportedCurrenciesMethodHelper(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void convert(hipstershop.Demo.CurrencyConversionRequest request,
        io.grpc.stub.StreamObserver<hipstershop.Demo.Money> responseObserver) {
      asyncUnaryCall(
          getChannel().newCall(getConvertMethodHelper(), getCallOptions()), request, responseObserver);
    }
  }

  /**
   */
  public static final class CurrencyServiceBlockingStub extends io.grpc.stub.AbstractStub<CurrencyServiceBlockingStub> {
    private CurrencyServiceBlockingStub(io.grpc.Channel channel) {
      super(channel);
    }

    private CurrencyServiceBlockingStub(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected CurrencyServiceBlockingStub build(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      return new CurrencyServiceBlockingStub(channel, callOptions);
    }

    /**
     */
    public hipstershop.Demo.GetSupportedCurrenciesResponse getSupportedCurrencies(hipstershop.Demo.GetSupportedCurrenciesRequest request) {
      return blockingUnaryCall(
          getChannel(), getGetSupportedCurrenciesMethodHelper(), getCallOptions(), request);
    }

    /**
     */
    public hipstershop.Demo.Money convert(hipstershop.Demo.CurrencyConversionRequest request) {
      return blockingUnaryCall(
          getChannel(), getConvertMethodHelper(), getCallOptions(), request);
    }
  }

  /**
   */
  public static final class CurrencyServiceFutureStub extends io.grpc.stub.AbstractStub<CurrencyServiceFutureStub> {
    private CurrencyServiceFutureStub(io.grpc.Channel channel) {
      super(channel);
    }

    private CurrencyServiceFutureStub(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected CurrencyServiceFutureStub build(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      return new CurrencyServiceFutureStub(channel, callOptions);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<hipstershop.Demo.GetSupportedCurrenciesResponse> getSupportedCurrencies(
        hipstershop.Demo.GetSupportedCurrenciesRequest request) {
      return futureUnaryCall(
          getChannel().newCall(getGetSupportedCurrenciesMethodHelper(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<hipstershop.Demo.Money> convert(
        hipstershop.Demo.CurrencyConversionRequest request) {
      return futureUnaryCall(
          getChannel().newCall(getConvertMethodHelper(), getCallOptions()), request);
    }
  }

  private static final int METHODID_GET_SUPPORTED_CURRENCIES = 0;
  private static final int METHODID_CONVERT = 1;

  private static final class MethodHandlers<Req, Resp> implements
      io.grpc.stub.ServerCalls.UnaryMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ServerStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ClientStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.BidiStreamingMethod<Req, Resp> {
    private final CurrencyServiceImplBase serviceImpl;
    private final int methodId;

    MethodHandlers(CurrencyServiceImplBase serviceImpl, int methodId) {
      this.serviceImpl = serviceImpl;
      this.methodId = methodId;
    }

    @java.lang.Override
    @java.lang.SuppressWarnings("unchecked")
    public void invoke(Req request, io.grpc.stub.StreamObserver<Resp> responseObserver) {
      switch (methodId) {
        case METHODID_GET_SUPPORTED_CURRENCIES:
          serviceImpl.getSupportedCurrencies((hipstershop.Demo.GetSupportedCurrenciesRequest) request,
              (io.grpc.stub.StreamObserver<hipstershop.Demo.GetSupportedCurrenciesResponse>) responseObserver);
          break;
        case METHODID_CONVERT:
          serviceImpl.convert((hipstershop.Demo.CurrencyConversionRequest) request,
              (io.grpc.stub.StreamObserver<hipstershop.Demo.Money>) responseObserver);
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

  private static abstract class CurrencyServiceBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoFileDescriptorSupplier, io.grpc.protobuf.ProtoServiceDescriptorSupplier {
    CurrencyServiceBaseDescriptorSupplier() {}

    @java.lang.Override
    public com.google.protobuf.Descriptors.FileDescriptor getFileDescriptor() {
      return hipstershop.Demo.getDescriptor();
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.ServiceDescriptor getServiceDescriptor() {
      return getFileDescriptor().findServiceByName("CurrencyService");
    }
  }

  private static final class CurrencyServiceFileDescriptorSupplier
      extends CurrencyServiceBaseDescriptorSupplier {
    CurrencyServiceFileDescriptorSupplier() {}
  }

  private static final class CurrencyServiceMethodDescriptorSupplier
      extends CurrencyServiceBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoMethodDescriptorSupplier {
    private final String methodName;

    CurrencyServiceMethodDescriptorSupplier(String methodName) {
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
      synchronized (CurrencyServiceGrpc.class) {
        result = serviceDescriptor;
        if (result == null) {
          serviceDescriptor = result = io.grpc.ServiceDescriptor.newBuilder(SERVICE_NAME)
              .setSchemaDescriptor(new CurrencyServiceFileDescriptorSupplier())
              .addMethod(getGetSupportedCurrenciesMethodHelper())
              .addMethod(getConvertMethodHelper())
              .build();
        }
      }
    }
    return result;
  }
}
