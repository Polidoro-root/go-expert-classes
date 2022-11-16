package main

import (
	"go-expert/mathematic"
)

func main() {
	s := mathematic.Sum(10, 20)

	println(s)

	println(mathematic.A)

	car := mathematic.Car{Brand: "Fiat"}

	println(car.Brand)
}
