package main
import "fmt"
func pow(x int,y int) int{
	z:=1
	for i:=0;i<y;i++{
		z*=x
	}
	return z
}

func main(){
	fmt.Println(pow(2,3))
	fmt.Println(pow(3,2))
	fmt.Println(pow(2,10))
}