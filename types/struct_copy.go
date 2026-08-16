package types

import "fmt"

type Person struct {
	name     string
	age      int
	accounts []int
}

func (p Person) PrintPerson() {
	fmt.Printf("Name: %s, Age: %d\n", p.name, p.age)
	fmt.Println("Accounts: ", p.accounts)
}

func TestCopy() {
	ac := []int{1, 2, 3, 4, 5}

	p := Person{
		name:     "Bob",
		age:      25,
		accounts: ac,
	}

	p.PrintPerson()

	pCopy := p
	p.accounts[0] = 7
	// since pCopy uses variable assignment, it creates a shallow copy
	// of the slice, so changing p, changes pCopy

	pCopy.age = 29
	pCopy.PrintPerson()

	pCopy.accounts[1] = 99
	p.PrintPerson()
	pCopy.PrintPerson()
	// this mutates both p's accounts and pCopy's accounts: [7 99 3 4 5]
}
