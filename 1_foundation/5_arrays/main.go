package main

import "fmt"

func main() {
	var myArray [3]int

	myArray[0] = 10
	myArray[1] = 20
	myArray[2] = 30

	fmt.Println(len(myArray) - 1)
	fmt.Println(myArray[len(myArray)-1])

	for index, value := range myArray {
		fmt.Printf("The value of index %d is %d\n", index, value)
	}
}
