package main

import (
	"context"
	"fmt"
	"learn_proto/rpc"

	"google.golang.org/grpc"
)

func main() {
	request := rpc.Request{
		Id: 23,
	}
	// grpc.WithInsecure
	// 要求tls证书验证，添加此选项，跳过验证
	conn, err := grpc.Dial(":54321", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	client := rpc.NewManagementClient(conn)
	response, err := client.QueryById(context.Background(), &request)
	if err != nil {
		panic(err)
	}
	fmt.Println(response.Name)
}
