package main

import "fmt"

func main() {
	score := 85

	if score >= 90 && score <= 100 {
		fmt.Println("A")
	} else if score >= 80 && score < 90 {
		fmt.Println("B")
	} else if score >= 70 && score < 80 {
		fmt.Println("C")
	} else if score >= 60 && score < 70 {
		fmt.Println("D")
	} else if score < 60 {
		fmt.Println("F")
	} else {
		fmt.Println("不正な値です")
	}
}
