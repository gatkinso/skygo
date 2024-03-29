// Generated by the gRPC C++ plugin.
// If you make any local change, they will be lost.
// source: skywire.proto
#ifndef GRPC_skywire_2eproto__INCLUDED
#define GRPC_skywire_2eproto__INCLUDED

#include "skywire.pb.h"

#include <functional>
#include <grpcpp/impl/codegen/async_generic_service.h>
#include <grpcpp/impl/codegen/async_stream.h>
#include <grpcpp/impl/codegen/async_unary_call.h>
#include <grpcpp/impl/codegen/client_callback.h>
#include <grpcpp/impl/codegen/client_context.h>
#include <grpcpp/impl/codegen/completion_queue.h>
#include <grpcpp/impl/codegen/method_handler.h>
#include <grpcpp/impl/codegen/proto_utils.h>
#include <grpcpp/impl/codegen/rpc_method.h>
#include <grpcpp/impl/codegen/server_callback.h>
#include <grpcpp/impl/codegen/server_context.h>
#include <grpcpp/impl/codegen/service_type.h>
#include <grpcpp/impl/codegen/status.h>
#include <grpcpp/impl/codegen/stub_options.h>
#include <grpcpp/impl/codegen/sync_stream.h>

namespace grpc_impl {
class CompletionQueue;
class ServerCompletionQueue;
class ServerContext;
}  // namespace grpc_impl

namespace grpc {
namespace experimental {
template <typename RequestT, typename ResponseT>
class MessageAllocator;
}  // namespace experimental
}  // namespace grpc

namespace skywire {

class Skywire final {
 public:
  static constexpr char const* service_full_name() {
    return "skywire.Skywire";
  }
  class StubInterface {
   public:
    virtual ~StubInterface() {}
    virtual ::grpc::Status Generate(::grpc::ClientContext* context, const ::skywire::SkywireRequest& request, ::skywire::SkywireResponse* response) = 0;
    std::unique_ptr< ::grpc::ClientAsyncResponseReaderInterface< ::skywire::SkywireResponse>> AsyncGenerate(::grpc::ClientContext* context, const ::skywire::SkywireRequest& request, ::grpc::CompletionQueue* cq) {
      return std::unique_ptr< ::grpc::ClientAsyncResponseReaderInterface< ::skywire::SkywireResponse>>(AsyncGenerateRaw(context, request, cq));
    }
    std::unique_ptr< ::grpc::ClientAsyncResponseReaderInterface< ::skywire::SkywireResponse>> PrepareAsyncGenerate(::grpc::ClientContext* context, const ::skywire::SkywireRequest& request, ::grpc::CompletionQueue* cq) {
      return std::unique_ptr< ::grpc::ClientAsyncResponseReaderInterface< ::skywire::SkywireResponse>>(PrepareAsyncGenerateRaw(context, request, cq));
    }
    class experimental_async_interface {
     public:
      virtual ~experimental_async_interface() {}
      virtual void Generate(::grpc::ClientContext* context, const ::skywire::SkywireRequest* request, ::skywire::SkywireResponse* response, std::function<void(::grpc::Status)>) = 0;
      virtual void Generate(::grpc::ClientContext* context, const ::grpc::ByteBuffer* request, ::skywire::SkywireResponse* response, std::function<void(::grpc::Status)>) = 0;
      virtual void Generate(::grpc::ClientContext* context, const ::skywire::SkywireRequest* request, ::skywire::SkywireResponse* response, ::grpc::experimental::ClientUnaryReactor* reactor) = 0;
      virtual void Generate(::grpc::ClientContext* context, const ::grpc::ByteBuffer* request, ::skywire::SkywireResponse* response, ::grpc::experimental::ClientUnaryReactor* reactor) = 0;
    };
    virtual class experimental_async_interface* experimental_async() { return nullptr; }
  private:
    virtual ::grpc::ClientAsyncResponseReaderInterface< ::skywire::SkywireResponse>* AsyncGenerateRaw(::grpc::ClientContext* context, const ::skywire::SkywireRequest& request, ::grpc::CompletionQueue* cq) = 0;
    virtual ::grpc::ClientAsyncResponseReaderInterface< ::skywire::SkywireResponse>* PrepareAsyncGenerateRaw(::grpc::ClientContext* context, const ::skywire::SkywireRequest& request, ::grpc::CompletionQueue* cq) = 0;
  };
  class Stub final : public StubInterface {
   public:
    Stub(const std::shared_ptr< ::grpc::ChannelInterface>& channel);
    ::grpc::Status Generate(::grpc::ClientContext* context, const ::skywire::SkywireRequest& request, ::skywire::SkywireResponse* response) override;
    std::unique_ptr< ::grpc::ClientAsyncResponseReader< ::skywire::SkywireResponse>> AsyncGenerate(::grpc::ClientContext* context, const ::skywire::SkywireRequest& request, ::grpc::CompletionQueue* cq) {
      return std::unique_ptr< ::grpc::ClientAsyncResponseReader< ::skywire::SkywireResponse>>(AsyncGenerateRaw(context, request, cq));
    }
    std::unique_ptr< ::grpc::ClientAsyncResponseReader< ::skywire::SkywireResponse>> PrepareAsyncGenerate(::grpc::ClientContext* context, const ::skywire::SkywireRequest& request, ::grpc::CompletionQueue* cq) {
      return std::unique_ptr< ::grpc::ClientAsyncResponseReader< ::skywire::SkywireResponse>>(PrepareAsyncGenerateRaw(context, request, cq));
    }
    class experimental_async final :
      public StubInterface::experimental_async_interface {
     public:
      void Generate(::grpc::ClientContext* context, const ::skywire::SkywireRequest* request, ::skywire::SkywireResponse* response, std::function<void(::grpc::Status)>) override;
      void Generate(::grpc::ClientContext* context, const ::grpc::ByteBuffer* request, ::skywire::SkywireResponse* response, std::function<void(::grpc::Status)>) override;
      void Generate(::grpc::ClientContext* context, const ::skywire::SkywireRequest* request, ::skywire::SkywireResponse* response, ::grpc::experimental::ClientUnaryReactor* reactor) override;
      void Generate(::grpc::ClientContext* context, const ::grpc::ByteBuffer* request, ::skywire::SkywireResponse* response, ::grpc::experimental::ClientUnaryReactor* reactor) override;
     private:
      friend class Stub;
      explicit experimental_async(Stub* stub): stub_(stub) { }
      Stub* stub() { return stub_; }
      Stub* stub_;
    };
    class experimental_async_interface* experimental_async() override { return &async_stub_; }

   private:
    std::shared_ptr< ::grpc::ChannelInterface> channel_;
    class experimental_async async_stub_{this};
    ::grpc::ClientAsyncResponseReader< ::skywire::SkywireResponse>* AsyncGenerateRaw(::grpc::ClientContext* context, const ::skywire::SkywireRequest& request, ::grpc::CompletionQueue* cq) override;
    ::grpc::ClientAsyncResponseReader< ::skywire::SkywireResponse>* PrepareAsyncGenerateRaw(::grpc::ClientContext* context, const ::skywire::SkywireRequest& request, ::grpc::CompletionQueue* cq) override;
    const ::grpc::internal::RpcMethod rpcmethod_Generate_;
  };
  static std::unique_ptr<Stub> NewStub(const std::shared_ptr< ::grpc::ChannelInterface>& channel, const ::grpc::StubOptions& options = ::grpc::StubOptions());

  class Service : public ::grpc::Service {
   public:
    Service();
    virtual ~Service();
    virtual ::grpc::Status Generate(::grpc::ServerContext* context, const ::skywire::SkywireRequest* request, ::skywire::SkywireResponse* response);
  };
  template <class BaseClass>
  class WithAsyncMethod_Generate : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service* /*service*/) {}
   public:
    WithAsyncMethod_Generate() {
      ::grpc::Service::MarkMethodAsync(0);
    }
    ~WithAsyncMethod_Generate() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status Generate(::grpc::ServerContext* /*context*/, const ::skywire::SkywireRequest* /*request*/, ::skywire::SkywireResponse* /*response*/) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    void RequestGenerate(::grpc::ServerContext* context, ::skywire::SkywireRequest* request, ::grpc::ServerAsyncResponseWriter< ::skywire::SkywireResponse>* response, ::grpc::CompletionQueue* new_call_cq, ::grpc::ServerCompletionQueue* notification_cq, void *tag) {
      ::grpc::Service::RequestAsyncUnary(0, context, request, response, new_call_cq, notification_cq, tag);
    }
  };
  typedef WithAsyncMethod_Generate<Service > AsyncService;
  template <class BaseClass>
  class ExperimentalWithCallbackMethod_Generate : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service* /*service*/) {}
   public:
    ExperimentalWithCallbackMethod_Generate() {
      ::grpc::Service::experimental().MarkMethodCallback(0,
        new ::grpc_impl::internal::CallbackUnaryHandler< ::skywire::SkywireRequest, ::skywire::SkywireResponse>(
          [this](::grpc::ServerContext* context,
                 const ::skywire::SkywireRequest* request,
                 ::skywire::SkywireResponse* response,
                 ::grpc::experimental::ServerCallbackRpcController* controller) {
                   return this->Generate(context, request, response, controller);
                 }));
    }
    void SetMessageAllocatorFor_Generate(
        ::grpc::experimental::MessageAllocator< ::skywire::SkywireRequest, ::skywire::SkywireResponse>* allocator) {
      static_cast<::grpc_impl::internal::CallbackUnaryHandler< ::skywire::SkywireRequest, ::skywire::SkywireResponse>*>(
          ::grpc::Service::experimental().GetHandler(0))
              ->SetMessageAllocator(allocator);
    }
    ~ExperimentalWithCallbackMethod_Generate() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status Generate(::grpc::ServerContext* /*context*/, const ::skywire::SkywireRequest* /*request*/, ::skywire::SkywireResponse* /*response*/) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    virtual void Generate(::grpc::ServerContext* /*context*/, const ::skywire::SkywireRequest* /*request*/, ::skywire::SkywireResponse* /*response*/, ::grpc::experimental::ServerCallbackRpcController* controller) { controller->Finish(::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "")); }
  };
  typedef ExperimentalWithCallbackMethod_Generate<Service > ExperimentalCallbackService;
  template <class BaseClass>
  class WithGenericMethod_Generate : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service* /*service*/) {}
   public:
    WithGenericMethod_Generate() {
      ::grpc::Service::MarkMethodGeneric(0);
    }
    ~WithGenericMethod_Generate() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status Generate(::grpc::ServerContext* /*context*/, const ::skywire::SkywireRequest* /*request*/, ::skywire::SkywireResponse* /*response*/) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
  };
  template <class BaseClass>
  class WithRawMethod_Generate : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service* /*service*/) {}
   public:
    WithRawMethod_Generate() {
      ::grpc::Service::MarkMethodRaw(0);
    }
    ~WithRawMethod_Generate() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status Generate(::grpc::ServerContext* /*context*/, const ::skywire::SkywireRequest* /*request*/, ::skywire::SkywireResponse* /*response*/) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    void RequestGenerate(::grpc::ServerContext* context, ::grpc::ByteBuffer* request, ::grpc::ServerAsyncResponseWriter< ::grpc::ByteBuffer>* response, ::grpc::CompletionQueue* new_call_cq, ::grpc::ServerCompletionQueue* notification_cq, void *tag) {
      ::grpc::Service::RequestAsyncUnary(0, context, request, response, new_call_cq, notification_cq, tag);
    }
  };
  template <class BaseClass>
  class ExperimentalWithRawCallbackMethod_Generate : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service* /*service*/) {}
   public:
    ExperimentalWithRawCallbackMethod_Generate() {
      ::grpc::Service::experimental().MarkMethodRawCallback(0,
        new ::grpc_impl::internal::CallbackUnaryHandler< ::grpc::ByteBuffer, ::grpc::ByteBuffer>(
          [this](::grpc::ServerContext* context,
                 const ::grpc::ByteBuffer* request,
                 ::grpc::ByteBuffer* response,
                 ::grpc::experimental::ServerCallbackRpcController* controller) {
                   this->Generate(context, request, response, controller);
                 }));
    }
    ~ExperimentalWithRawCallbackMethod_Generate() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status Generate(::grpc::ServerContext* /*context*/, const ::skywire::SkywireRequest* /*request*/, ::skywire::SkywireResponse* /*response*/) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    virtual void Generate(::grpc::ServerContext* /*context*/, const ::grpc::ByteBuffer* /*request*/, ::grpc::ByteBuffer* /*response*/, ::grpc::experimental::ServerCallbackRpcController* controller) { controller->Finish(::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "")); }
  };
  template <class BaseClass>
  class WithStreamedUnaryMethod_Generate : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service* /*service*/) {}
   public:
    WithStreamedUnaryMethod_Generate() {
      ::grpc::Service::MarkMethodStreamed(0,
        new ::grpc::internal::StreamedUnaryHandler< ::skywire::SkywireRequest, ::skywire::SkywireResponse>(std::bind(&WithStreamedUnaryMethod_Generate<BaseClass>::StreamedGenerate, this, std::placeholders::_1, std::placeholders::_2)));
    }
    ~WithStreamedUnaryMethod_Generate() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable regular version of this method
    ::grpc::Status Generate(::grpc::ServerContext* /*context*/, const ::skywire::SkywireRequest* /*request*/, ::skywire::SkywireResponse* /*response*/) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    // replace default version of method with streamed unary
    virtual ::grpc::Status StreamedGenerate(::grpc::ServerContext* context, ::grpc::ServerUnaryStreamer< ::skywire::SkywireRequest,::skywire::SkywireResponse>* server_unary_streamer) = 0;
  };
  typedef WithStreamedUnaryMethod_Generate<Service > StreamedUnaryService;
  typedef Service SplitStreamedService;
  typedef WithStreamedUnaryMethod_Generate<Service > StreamedService;
};

}  // namespace skywire


#endif  // GRPC_skywire_2eproto__INCLUDED
