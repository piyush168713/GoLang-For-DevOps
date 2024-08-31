package main

import (
	"02-GoLang-Basics-For-DevOps/shapes"
	"fmt"
)

func main(){
	radius := 2.0
	areaOfCircle := shapes.AreaOfCircle(radius)
	fmt.Printf("Area of Circle: %f\n", areaOfCircle)

	perimeterOfCircle := shapes.DiameterOfCircle(radius)
	fmt.Printf("Perimeter of Circle: %f\n", perimeterOfCircle)

	side := 6.0

	areaOfSquare := shapes.AreaOfSquare(side)
	fmt.Printf("Area of square: %f\n", areaOfSquare)

	perimeterOfSquare := shapes.PerimeterOfSquare(side)
	fmt.Printf("Perimeter of square: %f\n", perimeterOfSquare)

}