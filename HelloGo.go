package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello Go!")

	fmt.Println(sum(20, 30))

	fmt.Println(subs(20, 10))

	fmt.Println(abs(-1002))
}

func sum(i, j int) int {
	return i + j
}

func subs(i, j int) int {
	return i - j
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
