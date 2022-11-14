package main

import (
	"errors"
	"fmt"
)

func main() {
	value, err := sum(47, 2)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(value)
}

func sum(a, b int) (int, error) {
	result := a + b

	if result >= 50 {
		return 0, errors.New("The sum is greater than 50")
	}

	return result, nil
}
