package main

import "fmt"

type Person struct {
	// 名前
	Name string
	// 年齢
	Age int
}

func main() {
	// Person型のスライスを作成
	persons := []Person{
		// Alice, 20
		{Name: "Alice", Age: 20},
		// Bob, 25
		{Name: "Bob", Age: 25},
		// Charlie, 30
		{Name: "Charlie", Age: 30},
	}
	// スライスの要素を1つずつ取り出して、名前と年齢を出力

	for _, p := range persons {
		fmt.Println(p.Name, p.Age)
	}
}
