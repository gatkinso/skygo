# skygo

# Generate PB

Go: protoc -Iproto --go_out=plugins=grpc:skywire proto/skywire.proto

C++: protoc -Iproto --cpp_out=c++/client++ proto/skywire.proto
     protoc -Iproto --grpc_out=. --plugin=protoc-gen-grpc=$HOME/apps/grpc-debug/bin/grpc_cpp_plugin proto/skywire.proto
