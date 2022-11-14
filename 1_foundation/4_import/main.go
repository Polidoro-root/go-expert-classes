package main

import "fmt"

const f = "CONSTANT"

type ID int

var (
	a bool    = true
	b int     = 10
	c string  = "joao"
	d float64 = 1.2
	h ID      = 1
)

func main() {

	fmt.Printf("The type of 'h' is %T", h)
}
