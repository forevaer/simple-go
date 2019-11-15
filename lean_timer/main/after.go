package main

import "time"

func main() {
	t := time.NewTicker(3 * time.Second)
	for {
		select {
		case <-t.C:
			println(time.Now().Format("15:04:05"))
		default:

		}
	}
}
