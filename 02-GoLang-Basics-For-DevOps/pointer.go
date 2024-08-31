package main

import "fmt"

func main(){
	var a int = 10
	var b *int = &a

	fmt.Println("a:", a)
	fmt.Println("b:", b)
	fmt.Println("*b:", *b)
}
