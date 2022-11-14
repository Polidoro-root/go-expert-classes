package main

import "fmt"

func main() {
	salaries := map[string]int{"Joao": 1000, "Matheus": 2000}

	delete(salaries, "Joao")

	salaries["Pedro"] = 1500

	for name, salary := range salaries {
		fmt.Printf("The salary of %s is %d\n", name, salary)
	}

	for _, salary := range salaries {
		fmt.Printf("The salary %d\n", salary)
	}

	// salary := make(map[string]int)
}
