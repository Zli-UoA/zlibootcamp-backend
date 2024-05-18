package main

import (
	"fmt"
)

func pow(x int, y int) int {
    return for y ; y>0 ;y--{
		x = x*x
	}
}

func main() {
    fmt.Println(pow(3, 5)) // 8
}