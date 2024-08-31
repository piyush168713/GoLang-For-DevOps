package main

import (
	"fmt"
)

func main(){
	fruits := map[string]int{"apple": 5, "banana": 7}

	fmt.Println("Length of map:", len(fruits))

	// Add elements
	fruits["orange"] = 10
	fmt.Println("Added 'orange':", fruits)

	// Retreive element
	applePrice := fruits["apple"]
	fmt.Println("Price of apple:", applePrice)

	// checking existence
	// value, ok := mapVariable[key]
	price, exists := fruits["kiwi"]
	if exists {
		fmt.Println("Price of kiwi:", price)
	} else {
		fmt.Println("Kiwi does not exist in the map")
	}

	orangePrice, exists := fruits["orange"]
	if exists {
		fmt.Println("Price of orange:", orangePrice)
	} else {
		fmt.Println("Orange does not exist in the map")
	}

	// Deleting elements
	delete(fruits, "banana");
	fmt.Println("After deleting 'banana':", fruits)
}