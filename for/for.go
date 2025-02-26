package main

import "fmt"

func main() {
	for i:=0; i<=3; i++ {
		i = i+1
		fmt.Println(i)		
	}

	for j:=0; j<3; j++ {
		fmt.Println(j)
	}

	for i := range make([]int, 3) {
		fmt.Println("range", i)
	}
}