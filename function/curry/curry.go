package main

import (
	"fmt"
	"math"
)

func main() {
	c := cal(4)
	fmt.Println(c.do(add(1)))
	fmt.Println(c.do(del(1)))
	fmt.Println(c.do(sqrt))
}

type cal float64
type opfunc func(b float64) float64

func (c cal) do(op opfunc) cal {
	return cal(op(float64(c)))
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
