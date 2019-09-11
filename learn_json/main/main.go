package main

import (
	"encoding/json"
	"fmt"
	"learn_json/person"
)

func main() {
	p := person.Person{
		Name:   "Godme",
		Age:    99,
		Gender: "male",
	}

	bs := parse(p)
	fmt.Println(string(bs))

	var np person.Person
	format(bs, &np)
	fmt.Println(string(parse(np)))

}

func parse(source interface{}) []byte {
	result, _ := json.Marshal(source)
	return result
}

func format(source []byte, target interface{}) interface{} {
	_ = json.Unmarshal(source, target)
	return target
}
