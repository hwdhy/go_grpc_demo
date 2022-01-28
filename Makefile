gen:
	protoc --proto_path=proto proto/*.proto --go_out=plugins=grpc:pb

gens:
	protoc --proto_path=proto -I=D:\dev\GO\src  proto/*.proto --go_out=plugins=grpc:pb

clean:
	rm D:\person\code\grpc_project\pb\*.go

# -cover：覆盖率 -race：检测数据竞争
test:
	go test -cover -race ./...

server:
	go run ./cmd/server/main.go -port 8080

client:
	go run ./cmd/client/main.go -address 0.0.0.0:8080

cert:
	cd cert && gen.sh && cd ..

.PHONY: gens clean test server client cert