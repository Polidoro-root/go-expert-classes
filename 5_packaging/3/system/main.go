package main

import (
	"github.com/Polidoro-root/go-expert-classes/5_packaging/3/math"
	"github.com/google/uuid"
)

func main() {
	m := math.NewMath(1, 2)

	println(m.Add())
	println(uuid.New().String())
}
