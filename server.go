package main

import (
	"fmt"
	"google.golang.org/grpc"
	"mohuani-sub/mygrpc"
	"net"
)

func main() {
	createServer()
}

func createServer() {
	server, err := net.Listen("tcp", ":9998")
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("listen failed...")
	}
	myService := mygrpc.MyService{}
	grpcServer := grpc.NewServer()

	mygrpc.RegisterSubServiceServer(grpcServer, &myService)
	_ = grpcServer.Serve(server)

}
