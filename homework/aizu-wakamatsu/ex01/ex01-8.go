package main

import "fmt"

func main() {
	fmt.Println(kube(3))
}

func kube(x int) int {
	for i := 1; i < 3; i++ {
		x *= x
	}
	return x
}
