package main

import (
	"container/list"
	"fmt"
)

func main() {
	a := list.New()
	for index := 0; index < 10; index++ {
		temp := a.PushBack(index)
		fmt.Println(temp.Value)
	}
	print("============")
	for iterator := a.Front(); iterator != nil; iterator = iterator.Next() {
		fmt.Println(iterator.Value)
	}
}
