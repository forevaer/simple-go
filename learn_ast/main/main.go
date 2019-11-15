package main

import (
	"fmt"
	"runtime"
)

func main() {
	func() {
		func() {
			func() {
				func() {
					var flag bool = true
					var file string
					var line int
					for index := 0; flag; index++ {
						_, file, line, flag = runtime.Caller(index)
						fmt.Printf("file :%s , line :%d \n", file, line)
					}
				}()
			}()
		}()
	}()

}
