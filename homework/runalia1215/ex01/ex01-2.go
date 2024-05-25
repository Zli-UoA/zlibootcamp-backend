package main
import "fmt"
import "strconv"
func main() {
    a := 3
    b := 5
    fmt.Println("a + b = " +strconv.Itoa (a+b)) // 8
    fmt.Println("a - b = " + strconv.Itoa(a-b)) // -2
    fmt.Println("a * b = " + strconv.Itoa(a*b)) // 15
    fmt.Println("a / b = " + strconv.Itoa(a/b)) // 0
    fmt.Println("a % b = " + strconv.Itoa(a%b)) // 3
    fmt.Println("aの2乗 = " + strconv.Itoa(a*a)) // 9
}