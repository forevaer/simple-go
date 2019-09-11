package common

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
)

func Receive(connect net.Conn) {
	if connect == nil {
		fmt.Println("connect is nil")
		return
	}
	defer connect.Close()
	reader := bufio.NewReader(connect)
	for {
		line, _ := reader.ReadString('\n')
		fmt.Print(line)
	}
}

func Send(connect net.Conn) {
	if connect == nil {
		return
	}
	reader := bufio.NewReader(os.Stdin)
	defer connect.Close()
	for {
		line, _ := reader.ReadString('\n')
		_, _ = connect.Write([]byte(line))
	}
}
func Replay(connect net.Conn) {
	if connect == nil {
		return
	}
	_, _ = io.Copy(connect, connect)
}
