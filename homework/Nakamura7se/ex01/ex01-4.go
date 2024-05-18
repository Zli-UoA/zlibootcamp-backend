package main

import "fmt"

func main() {
	score := 85
	switch {
	case score >= 90:
		fmt.Println("A")
	case 80 <= score:
		fmt.Println("B")
	case 70 <= score:
		fmt.Println("C")
	case 60 <= score:
		fmt.Println("D")
	case score < 60:
		fmt.Println("F")
	default:
		fmt.Println("不正な値です")
	}
}
