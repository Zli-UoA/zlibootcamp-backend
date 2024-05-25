package main
import(
	"fmt"
	"math"
)
func main(){
	x:=2.0
	z:=1.0
	y:=z
	q:=0
	for {
		z -= (z*z - x) / (2*z)
		if(math.Abs(z-y)<0.000000001){
		//if(y==z){
			break
		}
		y=z
		q++;
	}
	fmt.Println(z)
	fmt.Println(math.Sqrt(x))
	fmt.Println(q)
}