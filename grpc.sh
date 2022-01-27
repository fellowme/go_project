
mkdir rpc_service
cd pdfile&&protoc --go_out=plugins=grpc:../rpc_service role.proto
cd ../