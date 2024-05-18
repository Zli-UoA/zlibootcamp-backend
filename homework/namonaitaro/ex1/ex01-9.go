package main

import (
	"fmt"
)

func pow(x int, y int) int {
	A := 1
	for i := 0; i < y; i++ {
		A *= x
	}
	return A
}

func main() {
	fmt.Println(pow(3, 3))
}
