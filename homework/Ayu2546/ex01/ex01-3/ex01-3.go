package main

import "fmt"

func main() {
	x := 2354
	flag := 0

	switch x % 3 {
	case 0:
		fmt.Print("Fizz")
		flag = 1
	}

	switch x % 5 {
	case 0:
		fmt.Print("Buzz")
		flag = 1
	}

	switch flag {
	case 0:
		fmt.Print(x)
	}

	fmt.Println()
}

/*
func main() {
	x := 2354
	flag := 0

	if x % 3 == 0 {
		fmt.Print("Fizz")
		flag = 1
	}

	if x % 5 == 0 {
		fmt.Print("Buzz")
		flag = 1
	}

	if flag == 0 {
		fmt.Print(x)
	}

	fmt.Println()
}
*/
