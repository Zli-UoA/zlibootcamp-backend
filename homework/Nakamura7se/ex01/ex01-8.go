package main

import "fmt"

func f(x int) int {
	return x * x * x
}
func main() {
	fmt.Println(f(3))
}
