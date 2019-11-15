package main

import bs "encoding/base64"

func main() {
	baseString := "%e5%8f%8c%e8%89%b2%e7%90%83"
	result, err := bs.StdEncoding.DecodeString(baseString)
	if err != nil {
		panic(err)
	}
	print(result)
}
