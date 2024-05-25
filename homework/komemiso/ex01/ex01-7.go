package main

import (
	"fmt"
	"math"
)

func main() {
	var i int
	x := 2.0
	z := 10.0
	fmt.Println(math.Sqrt(x))
	fmt.Println("")
	for i = 0; i <= 10; i++ {
		z -= (z*z - x) / (2 * z)
		fmt.Println(z)

	}
}
