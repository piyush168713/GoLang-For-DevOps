package main

import "fmt"

func main(){
	fmt.Println("Function as variable:")
	var a func(int) int
	a = func(x int) int { return x*x }
	result := a(5)

	fmt.Println("Result:", result)
}