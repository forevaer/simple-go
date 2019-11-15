package main

import (
	"bufio"
	"fmt"
	"net/http"
)

func main() {
	request, _ := http.NewRequest("GET", "http://www.baidu.com/s?wd=god", nil)
	response, _ := http.DefaultClient.Do(request)
	defer response.Body.Close()
	reader := bufio.NewReader(response.Body)
	for line, err := reader.ReadString('\n'); err != nil; {
		fmt.Println(line)
	}
}
