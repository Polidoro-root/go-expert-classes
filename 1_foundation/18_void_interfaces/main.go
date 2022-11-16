package main

import "fmt"

type I interface {
}

func main() {

	var x interface{} = 10
	var y interface{} = "Hello World"

	showType(x)
	showType(y)
}

func showType(t interface{}) {
	fmt.Printf("The type is %T and value is %v", t, t)
}
