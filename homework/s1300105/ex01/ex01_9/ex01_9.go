package main

import "fmt"

func pow(x, y int) int {
	a := x
	for i := 0; i < (y - 1); i++ {
		x *= a
	}
	return x
}

func main() {
	fmt.Println(pow(2, 3))
	fmt.Println(pow(3, 2))
	fmt.Println(pow(2, 10))

}
