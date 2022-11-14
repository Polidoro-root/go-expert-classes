package main

import "fmt"

type Client struct {
	Name     string
	Age      int
	IsActive bool
}

func main() {
	joao := Client{Name: "Joao", Age: 19, IsActive: true}

	joao.IsActive = false

	fmt.Println(joao)

}
