package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("learn_args/main/main.go")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
