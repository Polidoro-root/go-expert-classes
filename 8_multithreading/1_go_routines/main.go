package main

import (
	"fmt"
	"time"
)

func task(channel chan string, name string) {
	for i := 0; i < 10; i++ {
		fmt.Printf("Task %s is running %d\n", name, i)
		time.Sleep(time.Millisecond * 500)
	}

	channel <- name
}

func main() {
	channel := make(chan string)

	go task(channel, "A")
	go task(channel, "B")
	go task(channel, "C")
	go task(channel, "D")
	go task(channel, "E")
	go task(channel, "F")
	go task(channel, "G")
	go task(channel, "H")
	go task(channel, "I")
	go task(channel, "J")

	select {
	case task := <-channel:
		fmt.Printf("Task %s done", task)
	}

}
