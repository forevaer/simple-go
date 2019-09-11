package main

import (
	"fmt"
	"learn_socket/common"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	exit := make(chan os.Signal, 1)
	conn, err := net.Dial("tcp", ":9090")
	if err != nil {
		panic(err)
	}
	//go common.Receive(conn)
	go common.Send(conn)
	go common.Receive(conn)
	signal.Notify(exit, syscall.SIGQUIT)
	<-exit
	fmt.Printf("end")
}
