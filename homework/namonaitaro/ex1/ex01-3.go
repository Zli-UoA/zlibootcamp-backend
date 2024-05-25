package main

import "fmt"

func main() {
	x := 2353
	if x%3 == 0 && x%5 != 0 {
		fmt.Println("Fizz")
	} else if x%3 != 0 && x%5 == 0 {
		fmt.Println("Bizz")
	} else if x%3 == 0 && x%5 == 0 {
		fmt.Println("FizzBizz")
	} else {
		fmt.Println(x)
	}
}
