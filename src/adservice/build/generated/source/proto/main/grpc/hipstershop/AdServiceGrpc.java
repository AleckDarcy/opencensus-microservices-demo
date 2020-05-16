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

import java.util.logging.Level;
import java.util.logging.Logger;

/**
 */
@javax.annotation.Generated(
    value = "by gRPC proto compiler (version 1.10.1)",
    comments = "Source: demo.proto")
public final class AdServiceGrpc {
  private static final Logger log = Logger.getLogger(AdServiceGrpc.class.getName());

  private AdServiceGrpc() {}

  public static final String SERVICE_NAME = "hipstershop.AdService";

  // Static method descriptors that strictly reflect the proto.
  @io.grpc.ExperimentalApi("https://github.com/grpc/grpc-java/issues/1901")
  @java.lang.Deprecated // Use {@link #getGetAdsMethod()} instead. 
  public static final io.grpc.MethodDescriptor<hipstershop.Demo.AdRequest,
      hipstershop.Demo.AdResponse> METHOD_GET_ADS = getGetAdsMethodHelper();

  private static volatile io.grpc.MethodDescriptor<hipstershop.Demo.AdRequest,
      hipstershop.Demo.AdResponse> getGetAdsMethod;

  @io.grpc.ExperimentalApi("https://github.com/grpc/grpc-java/issues/1901")
  public static io.grpc.MethodDescriptor<hipstershop.Demo.AdRequest,
      hipstershop.Demo.AdResponse> getGetAdsMethod() {
    return getGetAdsMethodHelper();
  }

  private static io.grpc.MethodDescriptor<hipstershop.Demo.AdRequest,
      hipstershop.Demo.AdResponse> getGetAdsMethodHelper() {
    io.grpc.MethodDescriptor<hipstershop.Demo.AdRequest, hipstershop.Demo.AdResponse> getGetAdsMethod;
    if ((getGetAdsMethod = AdServiceGrpc.getGetAdsMethod) == null) {
      synchronized (AdServiceGrpc.class) {
        if ((getGetAdsMethod = AdServiceGrpc.getGetAdsMethod) == null) {
          AdServiceGrpc.getGetAdsMethod = getGetAdsMethod = 
              io.grpc.MethodDescriptor.<hipstershop.Demo.AdRequest, hipstershop.Demo.AdResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(
                  "hipstershop.AdService", "GetAds"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  hipstershop.Demo.AdRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  hipstershop.Demo.AdResponse.getDefaultInstance()))
                  .setSchemaDescriptor(new AdServiceMethodDescriptorSupplier("GetAds"))
                  .build();
          }
        }
     }
     return getGetAdsMethod;
  }

  /**
   * Creates a new async stub that supports all call types for the service
   */
  public static AdServiceStub newStub(io.grpc.Channel channel) {
    return new AdServiceStub(channel);
  }

  /**
   * Creates a new blocking-style stub that supports unary and streaming output calls on the service
   */
  public static AdServiceBlockingStub newBlockingStub(
      io.grpc.Channel channel) {
    return new AdServiceBlockingStub(channel);
  }

  /**
   * Creates a new ListenableFuture-style stub that supports unary calls on the service
   */
  public static AdServiceFutureStub newFutureStub(
      io.grpc.Channel channel) {
    return new AdServiceFutureStub(channel);
  }

  /**
   */
  public static abstract class AdServiceImplBase implements io.grpc.BindableService {

    /**
     */
    public void getAds(hipstershop.Demo.AdRequest request,
        io.grpc.stub.StreamObserver<hipstershop.Demo.AdResponse> responseObserver) {
      asyncUnimplementedUnaryCall(getGetAdsMethodHelper(), responseObserver);
    }

    @java.lang.Override public final io.grpc.ServerServiceDefinition bindService() {
      return io.grpc.ServerServiceDefinition.builder(getServiceDescriptor())
          .addMethod(
            getGetAdsMethodHelper(),
            asyncUnaryCall(
              new MethodHandlers<
                hipstershop.Demo.AdRequest,
                hipstershop.Demo.AdResponse>(
                  this, METHODID_GET_ADS)))
          .build();
    }
  }

  /**
   */
  public static final class AdServiceStub extends io.grpc.stub.AbstractStub<AdServiceStub> {
    private AdServiceStub(io.grpc.Channel channel) {
      super(channel);
    }

    private AdServiceStub(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected AdServiceStub build(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      return new AdServiceStub(channel, callOptions);
    }

    /**
     */
    public void getAds(hipstershop.Demo.AdRequest request,
        io.grpc.stub.StreamObserver<hipstershop.Demo.AdResponse> responseObserver) {
      asyncUnaryCall(
          getChannel().newCall(getGetAdsMethodHelper(), getCallOptions()), request, responseObserver);
    }
  }

  /**
   */
  public static final class AdServiceBlockingStub extends io.grpc.stub.AbstractStub<AdServiceBlockingStub> {
    private AdServiceBlockingStub(io.grpc.Channel channel) {
      super(channel);
    }

    private AdServiceBlockingStub(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected AdServiceBlockingStub build(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      return new AdServiceBlockingStub(channel, callOptions);
    }

    /**
     */
    public hipstershop.Demo.AdResponse getAds(hipstershop.Demo.AdRequest request) {
      return blockingUnaryCall(
          getChannel(), getGetAdsMethodHelper(), getCallOptions(), request);
    }
  }

  /**
   */
  public static final class AdServiceFutureStub extends io.grpc.stub.AbstractStub<AdServiceFutureStub> {
    private AdServiceFutureStub(io.grpc.Channel channel) {
      super(channel);
    }

    private AdServiceFutureStub(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected AdServiceFutureStub build(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      return new AdServiceFutureStub(channel, callOptions);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<hipstershop.Demo.AdResponse> getAds(
        hipstershop.Demo.AdRequest request) {
      return futureUnaryCall(
          getChannel().newCall(getGetAdsMethodHelper(), getCallOptions()), request);
    }
  }

  private static final int METHODID_GET_ADS = 0;

  private static final class MethodHandlers<Req, Resp> implements
      io.grpc.stub.ServerCalls.UnaryMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ServerStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ClientStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.BidiStreamingMethod<Req, Resp> {
    private final AdServiceImplBase serviceImpl;
    private final int methodId;

    MethodHandlers(AdServiceImplBase serviceImpl, int methodId) {
      this.serviceImpl = serviceImpl;
      this.methodId = methodId;
    }

    @java.lang.Override
    @java.lang.SuppressWarnings("unchecked")
    public void invoke(Req request, io.grpc.stub.StreamObserver<Resp> responseObserver) {
      switch (methodId) {
        case METHODID_GET_ADS:
          serviceImpl.getAds((hipstershop.Demo.AdRequest) request,
              (io.grpc.stub.StreamObserver<hipstershop.Demo.AdResponse>) responseObserver);
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

  private static abstract class AdServiceBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoFileDescriptorSupplier, io.grpc.protobuf.ProtoServiceDescriptorSupplier {
    AdServiceBaseDescriptorSupplier() {}

    @java.lang.Override
    public com.google.protobuf.Descriptors.FileDescriptor getFileDescriptor() {
      return hipstershop.Demo.getDescriptor();
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.ServiceDescriptor getServiceDescriptor() {
      return getFileDescriptor().findServiceByName("AdService");
    }
  }

  private static final class AdServiceFileDescriptorSupplier
      extends AdServiceBaseDescriptorSupplier {
    AdServiceFileDescriptorSupplier() {}
  }

  private static final class AdServiceMethodDescriptorSupplier
      extends AdServiceBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoMethodDescriptorSupplier {
    private final String methodName;

    AdServiceMethodDescriptorSupplier(String methodName) {
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
      synchronized (AdServiceGrpc.class) {
        result = serviceDescriptor;
        if (result == null) {
          serviceDescriptor = result = io.grpc.ServiceDescriptor.newBuilder(SERVICE_NAME)
              .setSchemaDescriptor(new AdServiceFileDescriptorSupplier())
              .addMethod(getGetAdsMethodHelper())
              .build();
        }
      }
    }
    return result;
  }
}
