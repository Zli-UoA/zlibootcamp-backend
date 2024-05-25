package main

import (
	"fmt"
	"math"
)

func main() {
	x := 2.0
	z := 2232.0
	cnt := 0
	for {
		if cnt == 10 {
			break
		}
		z -= (z*z - x) / (2 * z)
		cnt++
	}
	fmt.Println(z)
	fmt.Println(math.Sqrt(x))
}
