package main

import (
	"fmt"
	"math"
)

func main() {
	x := 2.0
	z := 1.0
	threshold := 1e-10 // 停止条件の閾値
	prevZ := 0.0

	for {
		prevZ = z
		z -= (z*z - x) / (2 * z)
		fmt.Println(z)
		if math.Abs(z-prevZ) < threshold {
			break
		}
	}

	fmt.Println("近似値:", z)
	fmt.Println("math.Sqrt(x):", math.Sqrt(x))
}
