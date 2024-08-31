package main

import (
	"fmt"
)

// 1. Single method interface
type Speaker interface {
	Speak() string
}

type Dog struct{
	Name string
}

func (d Dog) Speak() string {
	return d.Name + " says woof!"
}

// 2. Embedding Interface
type Animal interface{
	Speaker
	Move() string
}

type Bird struct{
	Name string
}

func (b Bird) Speak() string {
	return b.Name + " says meow!"
}

func (b Bird) Move() string{
	return b.Name + " flies!"
}

// 3. Empty Interface
func printAnything(v interface{}){
	fmt.Println(v)
}

// 4. Interface with multiple methods
type Vehicle interface{
	Start() string
	Stop() string
}

type Car struct{
	Model string
}

func (c Car) Start() string{
	return c.Model + " car started!"
}

func (c Car) Stop() string {
	return c.Model + " car Stopped!"
}

// 5. Functional Interface (single method interface)
// similar to speaker interface

// 6. Interface as a Constraint (Generics)
func Describe[T Speaker] (t T) {
	fmt.Println(t.Speak())
}



func main(){
	dog := Dog{Name: "Buddy"}
	bird := Bird{Name: "Tweety"}
	car := Car{Model: "Nissan"}

	fmt.Println(dog.Speak())
	fmt.Println(bird.Speak())
	fmt.Println(bird.Move())

	printAnything("An empty interface can hold anything")
	printAnything(42)

	fmt.Println(car.Start())
	fmt.Println(car.Stop())

	Describe(dog)
	Describe(bird)
}



