package test

import (
	"log"
	"testing"

	"github.com/garyburd/redigo/redis"
)

func TestConn(t *testing.T) {
	connect, err := redis.Dial("tcp", "127.0.0.1:6379")
	defer connect.Close()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("test redis connect success")
}
