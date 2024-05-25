package main

import "fmt"

func main() {
	for x := 1; x <= 100; x++ {
		// ここにコードを書く
		switch true {
		case x%3 == 0:
			fmt.Println("Fizz")
		case x%5 == 0:
			fmt.Println("Buzz")
		case x%3 == 0 && x%5 == 0:
			fmt.Println("FizzBuzz")
		default:
			fmt.Println(x)
		}
	}
}
