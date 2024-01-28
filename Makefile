.PHONY: tidy
tidy:
	go mod tidy

.PHONY: proto
proto:
	protoc --go_out=. ./proto/merchant/type/*.proto
	protoc --go-grpc_out=. ./proto/merchant/*.proto

.PHONY: run
run:
	go run main.go
