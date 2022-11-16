package main

func main() {
	// common For
	for i := 0; i < 10; i++ {
		println(i)
	}

	numbers := []string{"1", "2", "3"}

	// Foreach
	for _, v := range numbers {
		println(v)
	}

	// While
	i := 0

	for i < 10 {
		println(i)
		i++
	}

	// Infinite loop
	/* for {

		println("Hello World")
	} */
}
