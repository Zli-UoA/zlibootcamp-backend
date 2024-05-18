package main

import "fmt"

func main() {
	// 0から30までの数字に対して「FizzBuzz」のルールを適用して出力
	for x := 0; x <= 30; x++ {
		if x%3 == 0 && x%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if x%3 == 0 {
			fmt.Println("Fizz")
		} else if x%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(x)
		}
	}
}
