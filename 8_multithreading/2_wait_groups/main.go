package main

import (
	"fmt"
	"sync"
	"time"
)

func task(name string, wg *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		fmt.Printf("Task %s is running %d\n", name, i)
		time.Sleep(time.Millisecond * 500)
		wg.Done()
	}
}

func main() {
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(30)

	go task("A", &waitGroup)
	go task("B", &waitGroup)
	go task("C", &waitGroup)

	waitGroup.Wait()
}
