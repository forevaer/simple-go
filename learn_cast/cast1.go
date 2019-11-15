package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

func main() {
	pwd, _ := os.Getwd()
	fmt.Println(pwd)
	bs, err := ioutil.ReadFile(path.Join(pwd, "/learn_cast/0.png"))
	if err != nil {
		panic(err.Error())
	}
	pic_str := base64.StdEncoding.EncodeToString(bs)
	file, _ := os.Create("godme.txt")
	defer file.Close()
	_, _ = file.WriteString(pic_str)

}
