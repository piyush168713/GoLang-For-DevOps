package main
import "fmt"

func main(){
	a := map[string]int{"Alice": 30, "Bob": 20} // map with string keys and int values
	fmt.Println("a:",a)
	fmt.Println("Alice:", a["Alice"])
	fmt.Println("Bob:", a["Bob"])
}