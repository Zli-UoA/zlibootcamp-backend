package main
import (
    "fmt"
    "math"
)

func main() {
    x := 2.0
    z := 1.0
    for d := 1.0; d*d > 1e-10; z -= d { // 条件式を書く
        d = (z*z - x) / (2*z)
    }
    fmt.Println(z)
    fmt.Println(math.Sqrt(x))
}





