package main

import (
	"fmt"
)

func main(){
	fruits := map[string]int{"apple": 5, "banana": 7}

	fruits["orange"] = 10

	// iterate over the map
	for key, value := range fruits {
		fmt.Printf("%s: %d\n", key, value)
	}
}


