package main

import "fmt"

func main(){
	num := [5]int{10,20,30,40,50}

	// Using for loop
	for i := 0; i < len(num); i++ {
		fmt.Printf("Index: %d, Value: %d\n", i, num[i])
	}

	// Using a range loop
	for index, value := range num {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}
}