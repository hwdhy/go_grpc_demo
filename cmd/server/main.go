/*
@Time : 2022/1/16 10:30
@Author : Hwdhy
@File : main
@Software: GoLand
*/
package main

import (
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"grpc_project/pb"
	"grpc_project/service"
	"log"
	"net"
)

func main() {
	port := flag.Int("port", 0, "the server port")
	flag.Parse()
	log.Printf("start server on port %d\n", *port)

	laptopServer := service.NewLaptopServer(service.NewInmemoryLaptopStore())
	grpcServer := grpc.NewServer()

	pb.RegisterLaptopServiceServer(grpcServer, laptopServer)

	address := fmt.Sprintf("0.0.0.0:%d", *port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}
}
