package main
import "fmt"
func pow(x int, y int)int{
	result := 1
	for i := 0; i < y; i++{
		result *= x
	}
	return result
}

func main(){
	fmt.Println(pow(2, 3))
	fmt.Println(pow(3, 2))
	fmt.Println(pow(2, 10))
}
