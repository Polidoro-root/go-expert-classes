package main

import "fmt"

type I interface {
}

func main() {
	var var1 interface{} = "Joao Vitor"

	println(var1.(string))

	result, ok := var1.(int)

	fmt.Printf("The result is %v and ok is %v", result, ok)

}

func showType(t interface{}) {
	fmt.Printf("The type is %T and value is %v", t, t)
}
