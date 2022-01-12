gen:
	protoc --proto_path=proto proto/*.proto --go_out=plugins=grpc:pb

gens:
	protoc --proto_path=proto -I=D:\dev\GO\src  proto/*.proto --go_out=plugins=grpc:pb

clean:
	rm D:\person\code\grpc_project\pb\*.go

run:
	go run main.go