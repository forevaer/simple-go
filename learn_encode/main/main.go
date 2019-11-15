package main

import (
	"crypto/subtle"
	"encoding/base64"
	"fmt"
)

func main() {
	user := "godme"
	pass := "pass"
	str := []byte(fmt.Sprintf("%s:%s", user, pass))
	result := base64.StdEncoding.EncodeToString(str)
	fmt.Println(result)

	fmt.Println(subtle.ConstantTimeCompare([]byte("user"), []byte("user")))
}
