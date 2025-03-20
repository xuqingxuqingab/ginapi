// user service start
protoc \
  --proto_path=. \
  --go_out=. --go_opt=module=ginapi \
  --go-grpc_out=. --go-grpc_opt=module=ginapi \
  app/proto/user/*.proto

// order service start
protoc \
  --proto_path=. \
  --go_out=. --go_opt=module=ginapi \
  --go-grpc_out=. --go-grpc_opt=module=ginapi \
  app/proto/order/*.proto