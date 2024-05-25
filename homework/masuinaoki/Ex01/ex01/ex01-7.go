package main
import (
    "fmt"
    "math"
)
func main() {
    x := 2.0
    z := 1.0
		count:=0
    for i:=0; i<10; i++{ // 条件式を書く
        z -= (z*z - x) / (2*z)
				count++;
				if(z==math.Sqrt(x)){
					break
				}
    }
    fmt.Println(z)
    fmt.Println(math.Sqrt(x))
		fmt.Println("回数 :", count)
}