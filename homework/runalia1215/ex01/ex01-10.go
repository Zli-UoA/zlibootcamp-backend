package main
import "fmt"

type Person struct{
	Name string
	Age int
}

func main(){
	persons:=[]Person{
		{"Alice",20},{"Bob",25},{"Charlie",30},
	}
	for _,p := range persons{
		fmt.Println(p.Name,p.Age)
	}
}