package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	// Person型のスライスを作成
	persons := []Person{
		{"Alice", 20},
		{"Bob", 25},
		{"Charlie", 30},
		//Alice, 20
		//Bob, 25
		//Charlie, 30
	}
	// スライスの要素を1つずつ取り出して、名前と年齢を出力
	for _, p := range persons {
		fmt.Println(p.Name, p.Age)
	}
}
