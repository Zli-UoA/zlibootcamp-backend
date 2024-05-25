package main

import (
	"fmt"
	"math"
)

func main() {
	x := 2.0
	z := x
	p := 0
	for {
		new_z := z - (z*z-x)/(2*z)
		if math.Abs(new_z-z) < 1e-14 {
			break
		}
		p++
		z = new_z
	}
	fmt.Println(p)
	fmt.Println(z)
	fmt.Println(math.Sqrt(x))
}
