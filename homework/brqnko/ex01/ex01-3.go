package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 2354
	if x%15 == 0 {
		fmt.Println("FizzBuzz")
	} else if x%3 == 0 {
		fmt.Println("Fizz")
	} else if x%5 == 0 {
		fmt.Println("Buzz")
	} else {
		fmt.Println(strconv.Itoa(x))
	}
}