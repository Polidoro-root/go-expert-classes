package main

func main() {
	// memory -> address -> value
	a := 10

	var pointer *int = &a

	*pointer = 20

	b := &a

	*b = 30

	println(a)

}
