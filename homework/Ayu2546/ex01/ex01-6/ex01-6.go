package main

import "fmt"

func main() {
	for x := 1; x <= 100; x++ {
		if x%3 == 0 {
			fmt.Print("Fizz")
		}

		if x%5 == 0 {
			fmt.Print("Buzz")
		}

		if x%3 != 0 && x%5 != 0 {
			fmt.Print(x)
		}

		fmt.Println()
	}
}
