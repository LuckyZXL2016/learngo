package main

import (
	"fmt"
	"time"
)

// 此时channel，即可收又可发：func createWorker(id int) chan int
// 此时channel只可收， <-chan为此时只可发
func createWorker(id int) chan<- int {
	c := make(chan int)
	go func() {
		for {
			fmt.Printf("Worker %d received %c\n",
				id, <-c)
		}
	}()
	return c
}

func chanDemo() {
	var channels [10]chan<- int
	// 创建10个channel
	for i := 0; i < 10; i++ {
		channels[i] = createWorker(i)
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'A' + i
	}

	time.Sleep(time.Millisecond)
}

func main() {
	chanDemo()
}
