
# mkdir rpc_service
cd protobuf_file&&protoc --go_out=plugins=grpc:../rpc_service shop.proto
cd ../