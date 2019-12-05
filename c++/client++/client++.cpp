#include <grpcpp/grpcpp.h>
#include <grpc++/grpc++.h>
#include "skywire.grpc.pb.h"
#include "skywire.pb.h"

using grpc::Channel;
using grpc::ClientContext;
using grpc::Status;
using namespace skywire;

class SkywireClient
{
public:
    SkywireClient(std::shared_ptr<Channel> channel)
      : stub_(skywire::Skywire::NewStub(channel)) {}

    Status Generate(const skywire::SkywireRequest& req, skywire::SkywireResponse& res)
    {
        ClientContext context;

        Status status = stub_->Generate(&context, req, &res);

        if (!status.ok())
        {
        std::cout << status.error_code() << ": " << status.error_message()
                    << std::endl;
        }        

        return Status::OK;
    }

private:
    std::unique_ptr<skywire::Skywire::Stub> stub_;
};


int main(int argc, char** argv) {
  SkywireClient wirey(
      grpc::CreateChannel("localhost:1967", grpc::InsecureChannelCredentials())
  );

  skywire::SkywireRequest req;
  req.mutable_meta()->set_id(888);
  skywire::SkywireResponse res;
  auto status = wirey.Generate(req, res);

  std::cout << "Skywire received: " << res.DebugString() << std::endl;

  return 0;
}