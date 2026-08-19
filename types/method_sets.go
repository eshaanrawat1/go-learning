package types

import "fmt"

type Counter struct {
	n int
}

// value method modifes the copy
func (c Counter) ValueInc() {
	c.n++
}

// pointer method modifies the original
func (c *Counter) PtrInc() {
	c.n++
}

type VIncrementer interface{ ValueInc() }
type PIncrementer interface{ PtrInc() }

func TestInc() {
	c := Counter{0}

	c.ValueInc()
	fmt.Println("Counter value after value inc", c.n)
	// Counter value after value inc 0

	c.PtrInc()
	fmt.Println("Counter value after ptr inc", c.n)
	// Counter value after ptr inc 1

	c1 := Counter{10}

	var v VIncrementer = c1
	v.ValueInc()

	// var p PIncrementer = c1
	// p.PtrInc()
	// cannot use c1 (variable of struct type Counter) as PIncrementer value
	// in variable declaration: Counter does not implement PIncrementer
	// (method PtrInc has pointer receiver)

	var p PIncrementer = &c1
	p.PtrInc()
}
