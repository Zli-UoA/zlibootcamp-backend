package main

import "fmt"

func f(x int, y int) int {
	result := 1
	for i := 1; i <= y; i++ {
		result *= x
	}
	return result
}
func main() {
	fmt.Println(f(2, 3))
	fmt.Println(f(3, 2))
	fmt.Println(f(2, 10))
}
