package main

import (
	"fmt"
	"math"
)

func main() {
	a := 3
	b := 5
	fmt.Println(a + b)                   // a + bの結果を出力
	fmt.Println(a - b)                   //　a - bの結果を出力
	fmt.Println(a * b)                   // a * bの結果を出力
	fmt.Println(a / b)                   // a / bの結果を出力
	fmt.Println(a % b)                   // a / bの余りを出力
	fmt.Println(math.Pow(float64(a), 2)) // aの2乗の結果を出力
}
