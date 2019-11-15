package main

import (
	"encoding/base64"
	"io/ioutil"
	"os"
)

func main() {
	img_src, _ := ioutil.ReadFile("godme.txt")
	img_str := (string)(img_src)
	img_pic, _ := os.Create("godme.png")
	img_bs, _ := base64.StdEncoding.DecodeString(img_str)
	img_pic.Write(img_bs)
	img_pic.Close()
}
