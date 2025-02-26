package main

import (
	"fmt"
	"maps"
)

func main() {
	m := map[string]int{"foo": 1, "bar": 2}
	n := map[string]int{"bar": 2, "foo": 1}

	if maps.Equal(m,n) {
		fmt.Println("Maps are equal: ", m, n)
	} else {
		fmt.Println("Maps are not equal")
	}
}