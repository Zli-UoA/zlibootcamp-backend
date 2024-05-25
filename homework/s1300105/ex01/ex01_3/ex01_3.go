package main

import "fmt"

func main() {
	x := 2354

	if x%3 == 0 {
		fmt.Println("Fizz")
	} else if x%5 == 0 {
		fmt.Println("Buzz")
	} else if (x%3 == 0) && (x%5 == 0) {
		fmt.Println("FizzBuzz")
	} else {
		fmt.Println(x)
	}
}
