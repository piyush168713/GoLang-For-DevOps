package main

import (
	"fmt"
	"math"
)

// Interface - It is a set of methods
//           - A type implements an interface by implementing its methods.

type Shape interface{
	Area() float64
	Perimeter() float64
}

type Square struct{
	length float64
}

type Circle struct{
	radius float64
}

// For Square
func (s Square) Area() float64{
	return s.length*s.length
}

func (s Square) Perimeter() float64{
	return 4*s.length
}

// For Circle
func (c Circle) Area() float64{
	return math.Pi*c.radius*c.radius
}

func (c Circle) Perimeter() float64{
	return 2*math.Pi*c.radius
}

func printResult(s Shape){
	fmt.Printf("Area of %T is: %0.2f\n", s, s.Area())
	fmt.Printf("Perimeter of %T is: %0.2f\n", s, s.Perimeter())
}

func main(){
	fmt.Println("Interface:")
	
	var a Shape = Square{5}
	var b Shape = Circle{3}

	
	printResult(a);
	printResult(b);

	fmt.Println("a:",a)
	fmt.Println("b:",b)
}