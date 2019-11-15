package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

type Endpoint func(context.Context, interface{}) (interface{}, error)

func (endpoint Endpoint) exec(ctx context.Context, request interface{}) (interface{}, error) {
	return endpoint(ctx, request)
}

type MiddleWare func(Endpoint) Endpoint

func (middleware MiddleWare) wrapper(endpoint Endpoint) Endpoint {
	return middleware(endpoint)
}

type div_param struct {
	a int64
	b int64
}

func main() {
	div_endpoint := Endpoint(func(ctx context.Context, request interface{}) (interface{}, error) {
		if param, ok := request.(div_param); ok {
			return param.a / param.b, nil
		}
		return 0, fmt.Errorf("illegal param %v", request)
	})

	recover_ware := MiddleWare(func(endpoint Endpoint) Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			defer func() {
				if e := recover(); e != nil {
					err = errors.New("calc error")
				}
			}()
			return endpoint.exec(ctx, request)
		}
	})
	zero_ware := MiddleWare(func(endpoint Endpoint) Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			if param, ok := request.(div_param); ok {
				if param.b == 0 {
					return 0, errors.New("param b is zero")
				}
				return endpoint.exec(ctx, request)
			}
			return 0, fmt.Errorf("illegal param2 %v", request)
		}
	})
	use_middle := MiddleWare(func(endpoint Endpoint) Endpoint {
		return func(ctx context.Context, request interface{}) (interface{}, error) {
			start := time.Now()
			defer func() {
				fmt.Printf("use : %v \n", time.Since(start).Seconds())
			}()
			return endpoint.exec(ctx, request)
		}
	})
	final_endpoint := use_middle(zero_ware(recover_ware(div_endpoint)))
	param := div_param{
		a: 1,
		b: 0,
	}
	result, err := final_endpoint.exec(context.Background(), param)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}

func Chain(wares ...MiddleWare) MiddleWare {
	return func(endpoint Endpoint) Endpoint {
		return func(ctx context.Context, request interface{}) (interface{}, error) {
			for index := len(wares) - 1; index >= 0; index-- {
				endpoint = wares[index].wrapper(endpoint)
			}
			return endpoint(ctx, request)
		}

	}
}
