package main

import "fmt"

func cube(x int) int {
	return x * x * x
}

func main() {
	x := 3

	fmt.Println(cube(x))
}
