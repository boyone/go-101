package main

import "fmt"

func main() {
	number := 3

	scope := func() int {
		return number
	}

	fmt.Println(scope())

	number++

	fmt.Println(scope())
}
