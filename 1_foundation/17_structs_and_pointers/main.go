package main

import "fmt"

type Customer struct {
	name string
}

type Account struct {
	balance int
}

func NewAccount() *Account {
	return &Account{balance: 0}
}

// COPY STRUCT AS VALUE
func (c Customer) walk() {
	c.name = "Joao Vitor"

	fmt.Printf("The customer %v walked\n", c.name)
}

// USE POINTER OF STRUCT
func (a *Account) simulate(value int) int {
	a.balance += value
	println(a.balance)
	return a.balance
}

func main() {
	// joao := Customer{name: "Joao"}

	// joao.walk()

	// fmt.Printf("The customer name is %v\n", joao.name)

	account := Account{
		balance: 0,
	}

	account.simulate(300)

	println(account.balance)
}
