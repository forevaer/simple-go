package main

import (
	"fmt"
	"strings"
)

func main() {
	source := "source"
	strings.Compare(source, "source")
	strings.ToLower(source)
	strings.Contains(source, "target")
	for _, s := range strings.Fields("a b c d e f a d") {
		fmt.Println(s)
	}
	for _, s := range strings.FieldsFunc(source, func(rrr ...rune) func(r rune) bool {
		return func(r rune) bool {
			for _, rr := range rrr {
				if r == rr {
					return true
				}
			}
			return false
		}
	}('s', 'r')) {
		fmt.Println(s)
	}
	strings.HasPrefix(source, "s")
	strings.HasSuffix(source, "e")
	strings.Index(source, "a")
	strings.ReplaceAll(source, "o", "0")
}
