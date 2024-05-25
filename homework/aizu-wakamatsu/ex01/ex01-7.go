package main

import (
	"fmt"
	"math"
)

func main() {
	x := 2.0
	z := 1.0
	az := 0.0
	c := 0
	for c = 1; c <= 10; c++ { // 条件式を書く
		az = z
		z -= (z*z - x) / (2 * z)
		// fmt.Println(z)
		if az == z {
			break
		}
	}
	fmt.Print("iter ｶｲｽｳ ")
	fmt.Println(c)
	fmt.Print("z ｱﾀｨｲ    ")
	fmt.Println(z)
	fmt.Print("ﾍｲﾎｳｺﾝ    ")
	fmt.Println(math.Sqrt(x))
}
