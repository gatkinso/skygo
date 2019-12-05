cmake_minimum_required(VERSION 3.15)

project(client++ CXX)

set(TARGETNAME client++)

set(protobuf_MODULE_COMPATIBLE TRUE)
find_package(Protobuf CONFIG REQUIRED)
message(STATUS "Using protobuf ${protobuf_VERSION}")

set(_PROTOBUF_LIBPROTOBUF protobuf::libprotobuf)
set(_PROTOBUF_PROTOC $<TARGET_FILE:protobuf::protoc>)

find_package(gRPC CONFIG REQUIRED)
message(STATUS "Using gRPC ${gRPC_VERSION}")

set(_GRPC_GRPCPP_UNSECURE gRPC::grpc++_unsecure)
set(_GRPC_CPP_PLUGIN_EXECUTABLE $<TARGET_FILE:gRPC::grpc_cpp_plugin>)

set(_SOURCES client++.cpp
             skywire.pb.h
             skywire.pb.cc
             skywire.grpc.pb.h
             skywire.grpc.pb.cc
)

add_executable(${TARGETNAME} ${_SOURCES})

target_link_libraries(${TARGETNAME} ${_GRPC_GRPCPP_UNSECURE} ${_PROTOBUF_LIBPROTOBUF})

