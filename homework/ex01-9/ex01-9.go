package main
import "fmt"

func pow(x,y int) int {
	ans:=1
	for i:=0; i<y; i++{
		ans*=x
	}
	return ans
}

func main(){
	fmt.Println(pow(2,3))
	fmt.Println(pow(3,2))
	fmt.Println(pow(2,10))
}