package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("yes")
		_, err := writer.Write([]byte("fuck !"))
		fmt.Println(err)
	})
	err := http.ListenAndServe(":80", nil)
	fmt.Println(err)
}
