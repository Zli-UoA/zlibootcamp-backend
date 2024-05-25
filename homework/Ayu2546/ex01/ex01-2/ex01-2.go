package main

import (
	"fmt"
	"strconv"
)

func main() {
	a := 3
	b := 5
	ans1 := a + b
	ans2 := a - b
	ans3 := a * b
	ans4 := a / b
	ans5 := a % b
	ans6 := a * a
	fmt.Println("a + b = " + strconv.Itoa(ans1)) // 8
	fmt.Println("a - b = " + strconv.Itoa(ans2)) // -2
	fmt.Println("a * b = " + strconv.Itoa(ans3)) // 15
	fmt.Println("a / b = " + strconv.Itoa(ans4)) // 0
	fmt.Println("a % b = " + strconv.Itoa(ans5)) // 3
	fmt.Println("aの2乗 = " + strconv.Itoa(ans6))  // 9
}
