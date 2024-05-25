package main

import "fmt"

func pow(x, y int) int {
	S := 1
	for i := 0; i < y; i++ {
		S *= x
	}
	return S
}
func main() {
	fmt.Println(pow(2, 3))
	fmt.Println(pow(3, 2))
	fmt.Println(pow(2, 10))
}
