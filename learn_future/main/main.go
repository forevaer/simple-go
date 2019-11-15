package main

import (
	"context"
	"fmt"
	"time"
)

type Result struct {
	result chan int
	a      int
	b      int
}

func (result *Result) Run() {
	time.Sleep(time.Second * 3)
	result.result <- result.a + result.b
}

func Add(_a, _b int) Result {
	result := Result{
		result: make(chan int, 0),
		a:      _a,
		b:      _b,
	}
	go result.Run()
	return result
}

func main() {
	ctx, cancel := context.WithCancel(context.TODO())
	go func(c context.Context, cel context.CancelFunc) {
		t := time.NewTicker(time.Second)
		for {
			select {
			case <-t.C:
				fmt.Println(time.Now().Format("15:04:05"))
			case <-ctx.Done():
				cel()
				break
			}
		}
	}(ctx, cancel)
	sum := Add(1, 5)
	fmt.Println(<-sum.result)
}
