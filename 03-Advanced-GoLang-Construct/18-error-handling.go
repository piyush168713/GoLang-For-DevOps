package main

import (
	"errors"
	"fmt"
)

type customError struct {
	msg string
}

func (e *customError) Error() string {
	return e.msg
}

func someFunction(flag bool) error {
	if !flag {
		// Return a custom error
		return &customError{"Custom error occured!"}
	}
	return nil
}

func safeFunction() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from Panic:", r)
		}
	}()
	// This would cause a panic
	panic("A panic occured")
}

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a/b, nil
}

func main(){
	// Simple error return
	result, err := divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}

	// Another example with valid division
	result1, err1 := divide(10, 2)
	if err1 != nil {
		fmt.Println("Error:", err1)
	} else {
		fmt.Println("Result:", result1)
	}

	// custom error
	err2 := someFunction(false)
	if err2 != nil {
		fmt.Println("Error:", err2)
	}

	// Using panic and recover
	safeFunction() 
}