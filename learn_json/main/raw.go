package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type c1 struct {
	cunt string
}

type c2 struct {
	cunt int
}

func main() {
	var bs json.RawMessage = []byte("{\"cunt\" : 1}")
	c := c1{}
	if err := json.Unmarshal(bs, &c); err == nil {
		fmt.Println("c")
		fmt.Println(json.Marshal(c))
	}
	cc := c2{}
	if err := json.Unmarshal(bs, &cc); err == nil {
		fmt.Println("cc")
		_ = json.NewEncoder(os.Stdout).Encode(cc)
	}

}
