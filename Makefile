gen:
	protoc --proto_path=proto  proto/*.proto --go_out=plugins=grpc:pb

gens:
	protoc --proto_path=proto proto/*.proto --go_out=plugins=grpc:pb --grpc-gateway_out=:pb --swagger_out=:swagger

clean:
	rm pb/*

# -cover：覆盖率 -race：检测数据竞争
test:
	go test -cover -race ./...

server:
	go run ./cmd/server/main.go -port 8080

rest:
	go run ./cmd/server/main.go -port 8081 -type rest -endpoint 0.0.0.0:8080


server-tls:
	go run ./cmd/server/main.go -port 8080 -tls

server1:
	go run ./cmd/server/main.go -port 50051

server2:
	go run ./cmd/server/main.go -port 50052

server1-tls:
	go run ./cmd/server/main.go -port 50051 -tls

server2-tls:
	go run ./cmd/server/main.go -port 50052 -tls

client:
	go run ./cmd/client/main.go -address 0.0.0.0:8080

client-tls:
	go run ./cmd/client/main.go -address 0.0.0.0:8080 -tls

cert:
	cd cert && gen.sh && cd ..

.PHONY: gens clean test server client cert