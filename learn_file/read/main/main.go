package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("learn_file/read/main/go.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	_, _ = io.Copy(os.Stdout, file)
}
