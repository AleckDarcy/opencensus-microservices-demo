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
public final class EmailServiceGrpc {

  private EmailServiceGrpc() {}

  public static final String SERVICE_NAME = "hipstershop.EmailService";

  // Static method descriptors that strictly reflect the proto.
  @io.grpc.ExperimentalApi("https://github.com/grpc/grpc-java/issues/1901")
  @java.lang.Deprecated // Use {@link #getSendOrderConfirmationMethod()} instead. 
  public static final io.grpc.MethodDescriptor<hipstershop.Demo.SendOrderConfirmationRequest,
      hipstershop.Demo.Empty> METHOD_SEND_ORDER_CONFIRMATION = getSendOrderConfirmationMethodHelper();

  private static volatile io.grpc.MethodDescriptor<hipstershop.Demo.SendOrderConfirmationRequest,
      hipstershop.Demo.Empty> getSendOrderConfirmationMethod;

  @io.grpc.ExperimentalApi("https://github.com/grpc/grpc-java/issues/1901")
  public static io.grpc.MethodDescriptor<hipstershop.Demo.SendOrderConfirmationRequest,
      hipstershop.Demo.Empty> getSendOrderConfirmationMethod() {
    return getSendOrderConfirmationMethodHelper();
  }

  private static io.grpc.MethodDescriptor<hipstershop.Demo.SendOrderConfirmationRequest,
      hipstershop.Demo.Empty> getSendOrderConfirmationMethodHelper() {
    io.grpc.MethodDescriptor<hipstershop.Demo.SendOrderConfirmationRequest, hipstershop.Demo.Empty> getSendOrderConfirmationMethod;
    if ((getSendOrderConfirmationMethod = EmailServiceGrpc.getSendOrderConfirmationMethod) == null) {
      synchronized (EmailServiceGrpc.class) {
        if ((getSendOrderConfirmationMethod = EmailServiceGrpc.getSendOrderConfirmationMethod) == null) {
          EmailServiceGrpc.getSendOrderConfirmationMethod = getSendOrderConfirmationMethod = 
              io.grpc.MethodDescriptor.<hipstershop.Demo.SendOrderConfirmationRequest, hipstershop.Demo.Empty>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(
                  "hipstershop.EmailService", "SendOrderConfirmation"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  hipstershop.Demo.SendOrderConfirmationRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  hipstershop.Demo.Empty.getDefaultInstance()))
                  .setSchemaDescriptor(new EmailServiceMethodDescriptorSupplier("SendOrderConfirmation"))
                  .build();
          }
        }
     }
     return getSendOrderConfirmationMethod;
  }

  /**
   * Creates a new async stub that supports all call types for the service
   */
  public static EmailServiceStub newStub(io.grpc.Channel channel) {
    return new EmailServiceStub(channel);
  }

  /**
   * Creates a new blocking-style stub that supports unary and streaming output calls on the service
   */
  public static EmailServiceBlockingStub newBlockingStub(
      io.grpc.Channel channel) {
    return new EmailServiceBlockingStub(channel);
  }

  /**
   * Creates a new ListenableFuture-style stub that supports unary calls on the service
   */
  public static EmailServiceFutureStub newFutureStub(
      io.grpc.Channel channel) {
    return new EmailServiceFutureStub(channel);
  }

  /**
   */
  public static abstract class EmailServiceImplBase implements io.grpc.BindableService {

    /**
     */
    public void sendOrderConfirmation(hipstershop.Demo.SendOrderConfirmationRequest request,
        io.grpc.stub.StreamObserver<hipstershop.Demo.Empty> responseObserver) {
      asyncUnimplementedUnaryCall(getSendOrderConfirmationMethodHelper(), responseObserver);
    }

    @java.lang.Override public final io.grpc.ServerServiceDefinition bindService() {
      return io.grpc.ServerServiceDefinition.builder(getServiceDescriptor())
          .addMethod(
            getSendOrderConfirmationMethodHelper(),
            asyncUnaryCall(
              new MethodHandlers<
                hipstershop.Demo.SendOrderConfirmationRequest,
                hipstershop.Demo.Empty>(
                  this, METHODID_SEND_ORDER_CONFIRMATION)))
          .build();
    }
  }

  /**
   */
  public static final class EmailServiceStub extends io.grpc.stub.AbstractStub<EmailServiceStub> {
    private EmailServiceStub(io.grpc.Channel channel) {
      super(channel);
    }

    private EmailServiceStub(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected EmailServiceStub build(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      return new EmailServiceStub(channel, callOptions);
    }

    /**
     */
    public void sendOrderConfirmation(hipstershop.Demo.SendOrderConfirmationRequest request,
        io.grpc.stub.StreamObserver<hipstershop.Demo.Empty> responseObserver) {
      asyncUnaryCall(
          getChannel().newCall(getSendOrderConfirmationMethodHelper(), getCallOptions()), request, responseObserver);
    }
  }

  /**
   */
  public static final class EmailServiceBlockingStub extends io.grpc.stub.AbstractStub<EmailServiceBlockingStub> {
    private EmailServiceBlockingStub(io.grpc.Channel channel) {
      super(channel);
    }

    private EmailServiceBlockingStub(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected EmailServiceBlockingStub build(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      return new EmailServiceBlockingStub(channel, callOptions);
    }

    /**
     */
    public hipstershop.Demo.Empty sendOrderConfirmation(hipstershop.Demo.SendOrderConfirmationRequest request) {
      return blockingUnaryCall(
          getChannel(), getSendOrderConfirmationMethodHelper(), getCallOptions(), request);
    }
  }

  /**
   */
  public static final class EmailServiceFutureStub extends io.grpc.stub.AbstractStub<EmailServiceFutureStub> {
    private EmailServiceFutureStub(io.grpc.Channel channel) {
      super(channel);
    }

    private EmailServiceFutureStub(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected EmailServiceFutureStub build(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      return new EmailServiceFutureStub(channel, callOptions);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<hipstershop.Demo.Empty> sendOrderConfirmation(
        hipstershop.Demo.SendOrderConfirmationRequest request) {
      return futureUnaryCall(
          getChannel().newCall(getSendOrderConfirmationMethodHelper(), getCallOptions()), request);
    }
  }

  private static final int METHODID_SEND_ORDER_CONFIRMATION = 0;

  private static final class MethodHandlers<Req, Resp> implements
      io.grpc.stub.ServerCalls.UnaryMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ServerStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ClientStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.BidiStreamingMethod<Req, Resp> {
    private final EmailServiceImplBase serviceImpl;
    private final int methodId;

    MethodHandlers(EmailServiceImplBase serviceImpl, int methodId) {
      this.serviceImpl = serviceImpl;
      this.methodId = methodId;
    }

    @java.lang.Override
    @java.lang.SuppressWarnings("unchecked")
    public void invoke(Req request, io.grpc.stub.StreamObserver<Resp> responseObserver) {
      switch (methodId) {
        case METHODID_SEND_ORDER_CONFIRMATION:
          serviceImpl.sendOrderConfirmation((hipstershop.Demo.SendOrderConfirmationRequest) request,
              (io.grpc.stub.StreamObserver<hipstershop.Demo.Empty>) responseObserver);
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

  private static abstract class EmailServiceBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoFileDescriptorSupplier, io.grpc.protobuf.ProtoServiceDescriptorSupplier {
    EmailServiceBaseDescriptorSupplier() {}

    @java.lang.Override
    public com.google.protobuf.Descriptors.FileDescriptor getFileDescriptor() {
      return hipstershop.Demo.getDescriptor();
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.ServiceDescriptor getServiceDescriptor() {
      return getFileDescriptor().findServiceByName("EmailService");
    }
  }

  private static final class EmailServiceFileDescriptorSupplier
      extends EmailServiceBaseDescriptorSupplier {
    EmailServiceFileDescriptorSupplier() {}
  }

  private static final class EmailServiceMethodDescriptorSupplier
      extends EmailServiceBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoMethodDescriptorSupplier {
    private final String methodName;

    EmailServiceMethodDescriptorSupplier(String methodName) {
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
      synchronized (EmailServiceGrpc.class) {
        result = serviceDescriptor;
        if (result == null) {
          serviceDescriptor = result = io.grpc.ServiceDescriptor.newBuilder(SERVICE_NAME)
              .setSchemaDescriptor(new EmailServiceFileDescriptorSupplier())
              .addMethod(getSendOrderConfirmationMethodHelper())
              .build();
        }
      }
    }
    return result;
  }
}
