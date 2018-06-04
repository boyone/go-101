package main

import "fmt"

func main() {
	add := func(a, b int) int {
		return a + b
	}
	print(add, 2, 1)

	del := func(a, b int) int {
		return a - b
	}
	print(del, 2, 1)
}

type opfunc func(a, b int) int

func print(op opfunc, a, b int) {
	fmt.Println(op(a, b))
}