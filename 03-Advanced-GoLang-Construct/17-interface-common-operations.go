package main

import (
	"fmt"
)

type Greeter interface{
	Greet() string
}

type Person struct {
	Name string
}

func (p Person) Greet() string {
	return "Hello, " + p.Name
}

// Type Assertions
func printDetails(i interface{}) {
	str, ok := i.(string)
	if ok {
		fmt.Println("It's a string:", str)
	} else {
		fmt.Println("Not a string")
	}
}

// Interface with type switches
func Describe(i interface{}) {
	switch v := i.(type){
	case int:
		fmt.Println("Integer:", v)
	case string:
		fmt.Println("String:", v)
	default:
		fmt.Println("Unknown type")
	}
}

func main(){
	// Implementing a interface
	person := Person{Name: "Alice"}
	var greeter Greeter = person
	fmt.Println(greeter.Greet())

	// Type Assertions
	printDetails("Hello")
	printDetails(40)

	Describe("Hello World")
	Describe(2024)
	Describe(3.14)
}

