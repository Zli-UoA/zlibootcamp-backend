package main
import "fmt"

func main(){
	score := 85

	switch score {
	case score >= 90 && score <= 100:
		fmt.Println("A")
	}

}