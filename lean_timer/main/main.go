package main

import "time"

func main() {

	t := time.NewTimer(3 * time.Second)
	for {
		select {
		case <-t.C:
			{
				println("reset")
				// delay
				t.Reset(2 * time.Second)
			}
		default:

		}
	}
}
