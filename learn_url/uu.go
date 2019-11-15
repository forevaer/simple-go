package main

import (
	"net/url"
	"strings"
)

func main() {
	us := "http://localhost:8080/time?aaa=111&b=1212424"
	u, _ := url.Parse(us)
	println(u.Scheme)
	println(strings.TrimPrefix(us, "http://"))
}
