package main

import (
	"context"
	"fmt"
	"learn_proto/rpc"
	"net"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":54321")
	if err != nil {
		panic(err)
	}
	server := grpc.NewServer()
	rpc.RegisterManagementServer(server, &Data{})
	fmt.Println("register success")
	err = server.Serve(lis)
	fmt.Println(err)
	if err != nil {
		panic(err)
	}
	fmt.Println(err)
	fmt.Println("start successful")
}

type Data struct {
}

func (t *Data) QueryById(ctx context.Context, reqeust *rpc.Request) (response *rpc.Response, err error) {
	response = &rpc.Response{
		Name: fmt.Sprintf("uesr_%d", reqeust.Id),
	}
	return response, err
}
