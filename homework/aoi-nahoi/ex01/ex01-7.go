package main

import (
	"fmt"
	"math"
)

func main() {
	x := 2.0
	z := 1.0
	for i := 0; i < 10; i++ { // 条件式を書く
		z -= (z*z - x) / (2 * z)
	}
	fmt.Println(z)
	fmt.Println(math.Sqrt(x))
}
