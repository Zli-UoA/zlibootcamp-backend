package main

import "fmt"

func main() {
	x := 2354
	if x%3 == 0 {
		fmt.Println("Fizz")
	} else if x%5 == 0 {
		fmt.Println("Buzz")
	} else {
		fmt.Println("FizzBuzz")
	}
}
