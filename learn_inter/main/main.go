package main

import (
	"context"
	"fmt"
	"reflect"
)

type Endpoint func(ctx context.Context, request interface{}) (interface{}, error)

type Eps struct {
	WalkEP Endpoint
}

type walker struct {
}

func (w walker) Walk(msg string) (response interface{}, err error) {
	fmt.Println(msg)
	return
}

func Instance(instance interface{}, endpoints interface{}) interface{} {
	instanceValue := reflect.ValueOf(instance)
	instanceType := reflect.TypeOf(instance)
	endpointValue := reflect.ValueOf(endpoints)
	for index := 0; index < instanceType.NumMethod(); index++ {
		methodName := instanceType.Method(index).Name
		epFunc := endpointValue.Elem().FieldByName(methodName + "EP")
		if epFunc.IsValid() && epFunc.CanSet() {
			method := instanceValue.MethodByName(methodName)
			epFunc.Set(reflect.ValueOf(
				func(ctx context.Context, request interface{}) (response interface{}, err error) {
					res := method.Call([]reflect.Value{reflect.ValueOf(request.(string))})
					if res[0].IsValid() {
						response = res[0].Interface()
					}
					if res[1].IsValid() && res[1].Interface() != nil {
						err = res[1].Interface().(error)
					}
					return
				}))
		}
	}
	return endpointValue.Elem().Interface()
}

func main() {
	var param interface{} = "error"
	_, _ = Instance(walker{}, &Eps{}).(Eps).WalkEP(context.Background(), param)
}
