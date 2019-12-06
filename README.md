# skygo

# Generate PB

go get -u github.com/golang/protobuf/protoc-gen-go

Go: protoc -Iproto --go_out=plugins=grpc:proto proto/skywire.proto

C++: protoc -Iproto --cpp_out=c++/client++ proto/skywire.proto
     protoc -Iproto --grpc_out=. --plugin=protoc-gen-grpc=$HOME/apps/grpc-debug/bin/grpc_cpp_plugin proto/skywire.proto
     
Build grpc:
     cmake -DgRPC_CARES_PROVIDER="package" -DgRPC_PROTOBUF_PROVIDER="package" -DgRPC_SSL_PROVIDER="package" -DCMAKE_INSTALL_PREFIX="$HOME/apps/cares" -DgRPC_ZLIB_PROVIDER="package"  -DOPENSSL_ROOT_DIR=/usr/local/opt/openssl ..
