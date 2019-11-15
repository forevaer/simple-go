package main

import (
	"context"
	"fmt"
)

type EndPoint func(ctx context.Context, request interface{}) interface{}

type MiddleWare func(point EndPoint) EndPoint

func Chain(middleWares ...MiddleWare) MiddleWare {
	return func(point EndPoint) EndPoint {
		for index := len(middleWares) - 1; index >= 0; index-- {
			point = middleWares[index](point)
		}
		return point
	}
}

func createMiddleWare(msg string) MiddleWare {
	return func(point EndPoint) EndPoint {
		return func(ctx context.Context, request interface{}) interface{} {
			fmt.Println(msg, "PRE")
			defer func() {
				fmt.Println(msg, "POST")
			}()
			return point(ctx, request)
		}

	}
}

func main() {
	point := EndPoint(func(ctx context.Context, request interface{}) interface{} {
		fmt.Println("my point")
		return nil
	})
	Chain(createMiddleWare("first"), createMiddleWare("second"), createMiddleWare("third"))(point)(context.TODO(), struct {
	}{})
}
