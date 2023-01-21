package main

import (
	"fmt"
	"sync"
)

func main() {
	channel := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(10)

	go publish(channel)
	go reader(channel, &wg)
	wg.Wait()
}

func reader(ch chan int, wg *sync.WaitGroup) {
	for x := range ch {
		fmt.Printf("Received %d\n", x)
		wg.Done()
	}
}

func publish(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}

	close(ch)
}
