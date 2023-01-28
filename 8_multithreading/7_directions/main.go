package main

func publish(data string, ch chan<- string) {
	ch <- data
}

func read(ch <-chan string) {
	println(<-ch)
}

func main() {
	ch := make(chan string)

	go publish("Message", ch)
	read(ch)
}
