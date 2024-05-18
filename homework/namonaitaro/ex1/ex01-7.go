package main

import (
	"fmt"
	"math"
)

func main() {
	var x float64 = 5.0
	var z float64 = 3.0
	for i := 0; i <= 20; i++ {
		y := z
		z -= (z*z - x) / (2 * z)
		fmt.Println(z)
		if y-z <= 0.0000000000000001 || y-z <= -0.0000000000000001 {
			break
		}
	}
	if i <= 10 {
		fmt.Println("10より少ない　結果", math.Sqrt(x))
	} else {
		fmt.Println("10より少ない　結果", math.Sqrt(x))
	}
}
