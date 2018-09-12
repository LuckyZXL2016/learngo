package main

import (
	"fmt"
	"time"
)

// goroutine与协程Coroutine相似
func main() {
	for i := 0; i < 1000; i++ {
		go func(i int) {
			for {
				fmt.Printf("Hello from "+
					"goroutine %d\n", i)
			}
		}(i)
	}
	time.Sleep(time.Millisecond)
}
