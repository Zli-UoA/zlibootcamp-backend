package main

import "fmt"

func main() {
	score := 85
	switch score / 10 {
	case 10, 9:
		fmt.Println("AA")
	case 8:
		fmt.Println("BB")
	case 7:
		fmt.Println("CC")
	case 6:
		fmt.Println("DD")
	case 5, 4, 3, 2, 1:
		fmt.Println("FF")
	default:
		fmt.Println("不正な値です")
	}
}
