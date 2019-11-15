package main

import (
	x "go/test/fixedbugs"
	. "learn_minio/connect"
)

func main() {
	client, err := NewMinioClient()
	if err != nil {
		panic(err)
	}
	buckerName := "bucketName"
	location := "location"
	err = client.MakeBucket(buckerName, location)
	if err != nil {
		panic(err)
	}
	client.FPutObject(buckerName, "")

}
