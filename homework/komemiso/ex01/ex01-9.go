package main

import "fmt"

func pow(x, y int) int {
	i := 0
	j := 1

	for i = 0; i < y; i++ {
		j = j * x
	}
	return j
}
func main() {
	fmt.Println(pow(2, 10)) // world hello
}
