package main

import (
	"fmt"
	"math"
)

func main() {
	x := 2.0
	z := 1.0
	for d := 1.0; d*d > 1e-10; z -= d {
		z -= (z*z - x) / (2 * z)
	}
	fmt.Println(z)
	fmt.Println(math.Sqrt(x))
}

// :=を使って変数を宣言することもできます。var ~でもいける
// swicth文は　score >= 60 && score <= 50みたいなif文の書き方もいける
