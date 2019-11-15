package main

import (
	"fmt"
)

type String string

type Func func(string) interface{}

func (f Func) execute(params string) interface{} {
	return f(params)
}

func Show(content string) interface{} {
	fmt.Println(content)
	return nil
}
func main() {
	var a Func
	a = Show
	a.execute("fuck")
}
