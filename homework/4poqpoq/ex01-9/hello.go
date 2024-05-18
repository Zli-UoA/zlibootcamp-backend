package main

import "fmt"

// pow関数を定義
func pow(x, y float64) float64 {
	result := 1.0
	for i := 0; i < int(y); i++ {
		result *= x
	}
	return result
}

func main() {
	fmt.Println(pow(2, 3))  // 8
	fmt.Println(pow(3, 2))  // 9
	fmt.Println(pow(2, 10)) // 1024
}
