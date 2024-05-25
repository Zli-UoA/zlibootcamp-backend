package main

import "fmt"

func main() {
	x := 23541
	if x%3 == 0 {
		if x%5 == 0 {
			fmt.Println("FizzBuzz")
		} else {
			fmt.Println("Fizz")
		}
	} else if x%5 == 0 {
		fmt.Println("Buzz")
	} else {
		fmt.Println(x)
	}
	// :=を使って変数を宣言することもできます。var ~でもいける
}
