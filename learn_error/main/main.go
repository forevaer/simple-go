package main

import (
	"errors"
	"fmt"
)

var (
	Gerror = errors.New("gError")
)

func throw() {
	panic("throw")
}

func catch() {
	if err := recover(); err != nil {
		fmt.Println(err)
	}
}
func main() {
	//defer catch()
	//throw()

	for index := 5; index > -5; index-- {
		func(value int) {
			defer catch()
			fmt.Println(300 / index)
		}(index)
	}
}
