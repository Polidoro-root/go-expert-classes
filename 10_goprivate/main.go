package main

import (
	"fmt"

	"github.com/Polidoro-root/go-expert-classes/9_events/pkg/events"
)

func main() {
	ed := events.NewEventDispatcher()

	fmt.Println(ed)
}
