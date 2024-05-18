package main

import (
	"fmt"
	"strconv"
)

func main() {
	for num := 0; num <= 30; num++ {
		fmt.Print(strconv.Itoa(num) + "\t")
	}
	fmt.Println()
}
