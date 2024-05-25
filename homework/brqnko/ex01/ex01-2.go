package main

import (
	"fmt"
)

func main() {
	a := 3
	b := 5
	fmt.Println("a + b = ", a+b) // 8
	fmt.Println("a - b = ", a-b) // -2
	fmt.Println("a * b = ", a*b) // 15
	fmt.Println("a / b = ", a/b) // 0
	fmt.Println("a % b = ", a%b) // 3
	fmt.Println("aの2乗 = ", a*a)  // 9
	// fmt.Printf("a + b = %d\n", a+b)  // 8
	// fmt.Printf("a - b = %d\n", a-b)  // -2
	// fmt.Printf("a * b = %d\n", a*b)  // 15
	// fmt.Printf("a / b = %d\n", a/b)  // 0
	// fmt.Printf("a %% b = %d\n", a%b) // 3
	// fmt.Printf("aの2乗 = %d\n", a*a)   // 9
}
