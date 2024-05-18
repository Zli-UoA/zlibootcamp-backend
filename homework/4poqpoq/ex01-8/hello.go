package main

import "fmt"

// cube関数を定義
func cube(x float64) float64 {
	return x * x * x
}

func main() {
	result := cube(3)
	fmt.Println(result) // 27
}
