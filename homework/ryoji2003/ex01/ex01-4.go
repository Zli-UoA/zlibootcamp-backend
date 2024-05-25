package main
import "fmt"
func main() {
    score := 85
    switch {
		case score >= 90:
			fmt.Println("A")
		case score >= 80:
			fmt.Println("B")
		case score >= 70:
			fmt.Println("C")
		case score >= 60:
			fmt.Println("D")
		case score < 60:
			fmt.Println("F")
		case score < 0 || score > 100:
			fmt.Println("不正な値です")	
	}
}