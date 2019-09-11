package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	var buf bytes.Buffer
	msg := "there was some message"
	buf.WriteString(msg)
	apd := "\nthere are append message"
	_, _ = fmt.Fprint(&buf, apd)
	_, _ = buf.WriteTo(os.Stdout)
}
