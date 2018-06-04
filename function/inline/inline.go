package main

import (
	"fmt"
	"math"
)

func main() {
	c := myNumber{
		Num:  4,
		Add:  add(1),
		Sqrt: sqrt,
	}

	fmt.Println(c.do(c.Add))
	fmt.Println(c.do(c.Sqrt))

	c.Add = add(3)

	fmt.Println(c.do(c.Add))
}

type myNumber struct {
	Num  float64
	Add  opfunc
	Sqrt opfunc
}

type opfunc func(b float64) float64

func (c myNumber) do(op opfunc) float64 {
	return op(c.Num)
}

func add(b float64) func(a float64) float64 {
	return func(a float64) float64 {
		return a + b
	}
}

func del(b float64) func(a float64) float64 {
	return func(a float64) float64 {
		return a - b
	}
}

func sqrt(a float64) float64 {
	return math.Sqrt(a)
}
