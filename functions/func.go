package main

import (
	"fmt"
)
func plus(a int, b int) int {
	return a + b
}

func minor (a int, b int) int {
	return a - b
}

func main() {
	res1 := plus(1, 2)
	res2 := minor(1, 2)
	fmt.Println("1+2 =", res1)
	fmt.Println("1-2 =", res2)
}