package main

import "fmt"

func main() {
	score := 85
	switch {
	case score >= 90 && score <= 100:
		fmt.Println("A")
	case score >= 80 && score < 90:
		fmt.Println("B")
	case score >= 70 && score < 80:
		fmt.Println("C")
	case score >= 60 && score < 70:
		fmt.Println("D")
	case score < 60:
		fmt.Println("F")
	default:
		fmt.Println("不正な値です")
	}
}
