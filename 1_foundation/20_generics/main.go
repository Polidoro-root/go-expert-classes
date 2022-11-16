package main

// import "constraints"

type MyNumber int

type Number interface {
	~int | ~float64
}

func Sum[T Number](m map[string]T) T {
	var sum T

	for _, v := range m {
		sum += v
	}

	return sum
}

func Compare[T comparable](a, b T) bool {
	return a == b
}

func main() {
	m := map[string]int{"Wesley": 100, "Joao": 500, "Matheus": 250}

	m2 := map[string]float64{"Wesley": 100.5, "Joao": 500.1, "Matheus": 250.3}

	m3 := map[string]MyNumber{"Wesley": 100, "Joao": 500, "Matheus": 250}

	println(Sum(m))
	println(Sum(m2))
	println(Sum(m3))
	println(Compare(10, 2))
}
