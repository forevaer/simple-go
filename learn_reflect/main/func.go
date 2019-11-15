package main

import (
	"fmt"
	"os"
)

func main() {
	path, _ := os.Getwd()
	if err := os.Mkdir(fmt.Sprintf("%s/%s", path, `fuck\a`), 0777); err != nil {
		fmt.Println(err)
	}
}
