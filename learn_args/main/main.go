package main

import (
	"fmt"
	"os"
)

func main() {
	for index, param := range os.Args {
		fmt.Printf("%d. %s\n", index, param)
	}
}
