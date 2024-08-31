package main
import "fmt"

func main(){
	fmt.Println("Channel:")
	u := make(chan int)
	fmt.Println("u:",u)
}