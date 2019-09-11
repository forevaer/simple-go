package main

import (
	"fmt"
	"log"

	"github.com/garyburd/redigo/redis"
)

func main() {
	connect, err := redis.Dial("tcp", "127.0.0.1:6379")
	defer connect.Close()
	if err != nil {
		log.Fatalln(err)
	}
	for index := 0; index < 10; index++ {
		_, _ = connect.Do("lpush", "result", fmt.Sprintf("godme_%d", index))
	}

	result, _ := redis.Strings(connect.Do("lrange", "result", "0", "-1"))
	for index, valye := range result {
		fmt.Printf("key : %d\nvalue :%s\n", index, valye)
	}
}
