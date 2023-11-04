generate:
	protoc --proto_path=api api/*.proto --go_out=pkg/grpc_example
	protoc --proto_path=api api/*.proto --go-grpc_out=pkg/grpc_example
	protoc -I ./api --grpc-gateway_out ./pkg/grpc_example \
        --grpc-gateway_opt paths=source_relative \
        --grpc-gateway_opt generate_unbound_methods=true \
        api/grpc-example.proto