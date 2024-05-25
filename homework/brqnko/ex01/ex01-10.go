package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {
	// Person型のスライスを作成
	persons := []Person{
		Person{"Alice", 20},
		Person{"Bob", 25},
		Person{"Charlie", 30},
	}
	// スライスの要素を1つずつ取り出して、名前と年齢を出力
	for _, p := range persons {
		fmt.Println(p.name, p.age)
	}
}
