package main

import (
	"time"
)

func main() {
	for i := 0; i < 100000000; i++ {
		go func() {
			time.Sleep(time.Second * 100000)
		}()
	}
}
