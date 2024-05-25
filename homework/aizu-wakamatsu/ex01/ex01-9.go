package main

import "fmt"

func main() {
	fmt.Println(pow(2, 3))
	fmt.Println(pow(3, 2))
	fmt.Println(pow(2, 10))
}

func pow(x, y int) int {
	for i := 1; i < y; i++ {
		x *= x
	}
	return x
}
