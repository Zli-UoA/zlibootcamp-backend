package main

import "fmt"

func main() {
	for x := 1; x <= 100; x++ {
		fzbz(x)
	}
}

func fzbz(x int) {
	if (x%3) == 0 && (x%5) == 0 {
		fmt.Println("FizzBvzz")
	} else if (x % 3) == 0 {
		fmt.Println("Fizz")
	} else if (x % 5) == 0 {
		fmt.Println("Bvzz")
	} else {
		fmt.Println(x)
	}
}
