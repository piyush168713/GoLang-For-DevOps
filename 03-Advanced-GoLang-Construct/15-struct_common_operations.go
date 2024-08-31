package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string
	Age int
}

func (p *Person) sayHello(){
	fmt.Printf("Hello, my name is %s and I am %d years old.\n", p.Name, p.Age)
}

type Employee struct {
	ID int		  `json:"id"`
	Name string   `json:"name"`
}

func main(){
	alice := Person{Name: "Alice", Age: 22}
	fmt.Println("Struct Instance:", alice)

	fmt.Println("Name:", alice.Name)
	alice.Name = "Alice Smith"
	fmt.Println("Name:", alice.Name)

	// Pointer to struct
	bob := &Person{Name: "Bob", Age: 30}
	fmt.Println("Struct Instance:", bob)

	// Method to struct
	alice.sayHello()
	bob.sayHello()

	// Tagging struct field
	emp := Employee{ID: 1, Name: "John"}
	t := reflect.TypeOf(emp)
	field, _ := t.FieldByName("Name")
	fmt.Println("Tag on Name field:", field.Tag.Get("json"))
}