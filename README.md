# skygo

# Generate PB

go get -u github.com/golang/protobuf/protoc-gen-go

Go: protoc -Iproto --go_out=plugins=grpc:skywire proto/skywire.proto

C++: protoc -Iproto --cpp_out=c++/client++ proto/skywire.proto
     protoc -Iproto --grpc_out=. --plugin=protoc-gen-grpc=$HOME/apps/grpc-debug/bin/grpc_cpp_plugin proto/skywire.proto
