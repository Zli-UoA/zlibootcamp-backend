package main

import "fmt"

func main() {
	a := 3
	b := 5
	fmt.Println("a + b = ", a+b) // 8
	fmt.Println("a + b = ", b-a) // -2
	fmt.Println("a + b = ", b*3) // 15
	fmt.Println("a + b = ", b/b) // 0
	fmt.Println("a + b = ", a%b) // 3
	fmt.Println("a + b = ", a*a) // 9
}
