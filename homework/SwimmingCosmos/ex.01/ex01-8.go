package main

import "fmt"

func cube(x int) int {
	return x * x * x
}
func main() {
	fmt.Println(cube(3)) // 9
}

// :=を使って変数を宣言することもできます。var ~でもいける
// swicth文は　score >= 60 && score <= 50みたいなif文の書き方もいける
