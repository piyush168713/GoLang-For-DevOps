package main

import (
	"fmt"
)

func modifySlice(s []int){
	if len(s) > 0 {
		s[0] = 999
	}
	fmt.Println("Inside modifySlice, after modification:", s)
}

func modifyArray(s [5]int){
	if len(s) > 0 {
		s[0] = 999
	}
	fmt.Println("Inside modifySlice, after modification:", s)
}

func main(){
	// Slice
	originalSlice := []int{2,3,4,5,6}
	// Array
	originalArray := [5]int{2,3,4,5,6}

	// Before modify
	fmt.Println("Original Slice:", originalSlice)
	fmt.Println("Original Array:", originalArray)

	modifySlice(originalSlice)
	modifyArray(originalArray)

	// After modify
	fmt.Println("Original Slice:", originalSlice)
	fmt.Println("Original Array:", originalArray)
}


// Array works as pass by value
// Slice works as pass by reference