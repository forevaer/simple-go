package main

import (
	"fmt"
	"learn_reflect/person"
	"reflect"
)

func main() {
	p := person.NewPerson("godme", "male", 99)
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
