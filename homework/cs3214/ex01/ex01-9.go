package main
import "fmt"

func pow(x,y int) int{
	tmp := 1
	for i := 0; i < y; i++{
		tmp = tmp * x
	}
	return tmp
}


func main(){
	fmt.Println(pow(2, 3))
	fmt.Println(pow(3, 2))
	fmt.Println(pow(2, 10))
	
}