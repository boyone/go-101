package main

import "fmt"


func main() {
	add := func(a, b int) int {
		return a + b
	}
		
	fmt.Printf("address: %v\n", &add)

	print(add)
}

func print(f func(a, b int) int) {
	fmt.Printf("address: %v\n", f)
}