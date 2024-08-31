package main
import "fmt"

func main(){
	type Person struct{
		Name string
		Age int
	}

	var a Person = Person{"Piyush", 22}
	fmt.Println("a = ", a)
	fmt.Println("Name:", a.Name)
	fmt.Println("Age:", a.Age)
}