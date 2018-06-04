package main

import (
	"fmt"
	"math"
)

func main() {
	c := cal(4)
	add := func(a, b float64) float64 {
		return a + b
	}
	fmt.Println(c.do(add, 1))

	del := func(a, b float64) float64 {
		return a - b
	}
	fmt.Println(c.do(del, 1))

	sqrt := func(a, b float64) float64 {
		return math.Sqrt(a)
	}

	fmt.Println(c.do(sqrt, 0))
}

type cal float64
type opfunc func(a, b float64) float64

func (c cal) do(op opfunc, b float64) cal {
	return cal(op(float64(c), 1))
}