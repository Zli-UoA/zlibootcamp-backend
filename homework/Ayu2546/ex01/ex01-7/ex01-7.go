package main

import (
	"fmt"
	"math"
)

func main() {
	x := 2.0
	z := 1.0
	past := 0.0
	flag := 0

	for i := 0; i < 10; i++ {
		z -= ((z * z) - x) / (2 * z)

		if past == z {
			flag = 1
			break
		}

		past = z
	}

	if flag == 1 {
		fmt.Println("10回より少ない")
	}

	fmt.Println(z)
	fmt.Println(math.Sqrt(x))
}
