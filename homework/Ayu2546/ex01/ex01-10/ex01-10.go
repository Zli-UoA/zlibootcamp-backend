package main

import "fmt"

type Person struct {
	Name string //名前
	Age  int    //年齢
}

func main() {
	persons := []Person{
		{Name: "Alice", Age: 20},
		{Name: "Bob", Age: 25},
		{Name: "Charlie", Age: 30},
	}

	for _, p := range persons {
		fmt.Println(p.Name, p.Age)
	}
}
