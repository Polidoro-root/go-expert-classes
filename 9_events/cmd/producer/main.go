package main

import "github.com/Polidoro-root/go-expert-classes/9_events/pkg/rabbitmq"

func main() {
	ch, err := rabbitmq.OpenChannel()

	if err != nil {
		panic(err)
	}

	defer ch.Close()

	rabbitmq.Publish(ch, "amq.direct", "Hello World!")
}
