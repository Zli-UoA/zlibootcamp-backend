package main
import "fmt"
func main() {
    score := 85
    // ここにコードを書く
		var str string

		switch {
		case score>=90 && score>=100:
			str = "A"
		case score>=80:
			str = "B"
		case score>=70:
			str = "C"
		case score>=60:
			str = "D"
		case score<60:
			str = "F"
		default:
			str = "不正な値です"
		}

		fmt.Println(str)
}