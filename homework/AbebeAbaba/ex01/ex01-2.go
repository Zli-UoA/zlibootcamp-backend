package main

import (
	"fmt"
)

func mai() {
	a := 3
	b := 5
	fmt.Println("a + b = " + fmt.Sprint(a+b)) // 8
	fmt.Println("a - b = " + fmt.Sprint(a-b)) // -2
	fmt.Println("a * b = " + fmt.Sprint(a*b)) // 15
	fmt.Println("a / b = " + fmt.Sprint(a/b)) // 0
	fmt.Println("a % b = " + fmt.Sprint(a%b)) // 3
	fmt.Println("aの2乗 = " + fmt.Sprint(a*a))  // 9
}
