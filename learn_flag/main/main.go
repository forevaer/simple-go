package main

import (
	"encoding/json"
	"flag"
	"fmt"
)

const (
	username = "godme"
	password = "godme"
	template = "welcome %s\n"
	asshole = "asshole"

)
func main() {
	host := flag.String("h", "127.0.0.1", "host")
	port := flag.Int("p", 6379, "port")
	pass := flag.String("P", "password", "password")
	name := flag.String("u", "admin", "username")
	data := make(map[string]interface{}, 0)
	flag.Parse()
	data["host"] = host
	data["port"] = port
	data["user"] = name
	data["pass"] = pass
	connInfo,_ := json.MarshalIndent(data, "\t", "")
	fmt.Println(string(connInfo))
	if *name == username && *pass == password {
		hello(*name)
	}else {
		hello(asshole)
	}
}

func hello(name string){
	fmt.Printf(template, name)
}

