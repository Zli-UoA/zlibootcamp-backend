package main
import "fmt"
func main() {
    x := 2354
    // ここにコードを書く
		var str string
		if(x%3==0){
			str = "Fizz"
		}else if(x%5==0){
			str = "Buzz"
		}else if(x%3==0 && x%5==0){
			str = "FizzBuzz"
		}else{
			fmt.Println(x);
			return;
		}

		fmt.Println(str);
}