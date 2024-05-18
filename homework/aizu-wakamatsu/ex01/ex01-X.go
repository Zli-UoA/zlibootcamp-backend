package main

import "fmt"

type Person struct {
	Name string // 名前
	Age  int    // 年齢
}

func main() {
	// Person型のスライスを作成
	persons := []Person{
		{Name: "Alice", // Alice, 20
			Age: 20},
		{Name: "Gib", // Bob, 25
			Age: 25},
		{Name: "Charlie", // Charlie, 30
			Age: 30},
	}
	// スライスの要素を1つずつ取り出して、名前と年齢を出力
	for _, p := range persons {
		fmt.Println(p.Name, p.Age)
	}
}
