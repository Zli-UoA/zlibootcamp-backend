package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	// Person型のスライスを作成
	persons := []Person{
		{Name: "Alice", Age: 20},
		{Name: "Bob", Age: 25},
		{Name: "Charlie", Age: 30},
	}
	// スライスの要素を1つずつ取り出して、名前と年齢を出力
	for _, p := range persons {
		fmt.Println("Name:", p.Name, ", Age:", p.Age)
	}
}
