protobuf_generate:
	protoc --proto_path=../internal/proto --go_out=.. --go-grpc_out=.. service.proto
rebuild:
	docker-compose build
	docker-compose up