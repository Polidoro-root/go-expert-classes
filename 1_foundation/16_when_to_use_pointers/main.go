package main

func sum(a, b *int) int {
	*a = 50
	*b = 25

	return *a + *b
}

func main() {
	var1 := 1
	var2 := 2

	sum(&var1, &var2)

	println(var1)
	println(var2)
}
