package connect

import (
	"log"

	"github.com/minio/minio-go"
)

const (
	endpoint = "localhost:9000"
	username = "godme"
	password = "password"
	useSSL   = false
)

func NewMinioClient() (*minio.Client, error) {
	return minio.New(endpoint, username, password, useSSL)
}

func main() {
	// 初使化 minio client对象。
	minioClient, err := NewMinioClient()
	if err != nil {
		log.Fatalln(err)
	} else {
		println("fuck")
	}

	log.Printf("%#v\n", minioClient) // minioClient初使化成功
}
