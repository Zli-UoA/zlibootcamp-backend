package main

import "fmt"

func main() {
	score := 85
	switch score / 10 {
	case 10:
		fmt.Println("A")
	case 9:
		fmt.Println("A")
	case 8:
		fmt.Println("B")
	case 7:
		fmt.Println("C")
	case 6:
		fmt.Println("D")
	case 5:
		fmt.Println("F")
	case 4:
		fmt.Println("F")
	case 3:
		fmt.Println("F")
	case 2:
		fmt.Println("F")
	case 1:
		fmt.Println("F")
	case 0:
		fmt.Println("F")
	default:
		fmt.Println("不正な値です")
	}
	// :=を使って変数を宣言することもできます。var ~でもいける
	// swicth分は　score >= 60 && score <= 50の書き方もいける
}
