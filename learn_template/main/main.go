package main

import (
	"bytes"
	"fmt"
	"html/template"
	"strconv"
)

const tmpl = "{{ range .S}} {{.}} \n{{end}}"

func main() {
	buf := bytes.NewBuffer(nil)

	str := make([]string, 9)
	for index := 0; index < 10; index++ {
		str = append(str, strconv.Itoa(index))
	}
	template.Must(template.New("godme").Parse(tmpl)).Execute(buf, struct {
		S []string
	}{
		S: str,
	})

	fmt.Println(buf.String())
}
