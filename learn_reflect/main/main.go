package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name   string
	Age    int
	Gender string
}

func main() {
	p := Person{"godme", 99, "male"}
	// type
	t := reflect.TypeOf(p)
	// value
	v := reflect.ValueOf(p)
	length := t.NumField()
	for index := 0; index < length; index++ {
		field := t.Field(index)
		// value
		fieldValue := v.Field(index)
		// name
		fieldName := field.Name
		// type
		fieldType := field.Type
		fmt.Printf("%s %s = %v\n", fieldName, fieldType, fieldValue)
	}
}
