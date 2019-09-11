package main

import (
	"fmt"
	"learn_socket/common"
	"net"
)

func main() {

	server, err := net.Listen("tcp", ":9090")
	if err != nil {
		panic(err)
	}
	users := make(map[net.Addr]net.Conn, 0)
	for {
		conn, err := server.Accept()
		users[conn.RemoteAddr()] = conn
		fmt.Println("connect successful")
		if err == nil {
			go func() {
				defer conn.Close()
				common.Replay(conn)
			}()
		}
	}
}
