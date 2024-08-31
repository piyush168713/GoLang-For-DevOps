package main

import "fmt"

func main() {
	// 1. Using the var Keyword without Initializing
	var slice1 []int
	fmt.Println("Slice 1 (uninitialized):", slice1)

	// Initialize slice1 before using it
	slice1 = []int{1, 2, 3}
	fmt.Println("Slice 1 (initialized):", slice1)

	// 2. Using a Slice Literal
	slice2 := []int{4, 5, 6}
	fmt.Println("Slice 2:", slice2)

	slice3 := make([]int, 3)   // len and cap are both 3
	slice3[0] = 7
	slice3[1] = 8
	slice3[2] = 9

	fmt.Println("Slice 3:", slice3)


	// 3. Using the make Function
	slice4 := make([]int, 3, 5)   // len is 3 & cap is 5
	slice4[0] = 10
	slice4[1] = 12
	slice4[2] = 13

	fmt.Println("Slice 4:", slice4)
	fmt.Printf("Length: %d, Capacity: %d\n", len(slice4), cap(slice4))

	// slice4[3] = 15   // it do not work coz len is 3
	// fmt.Println("Slice 4:", slice4)

	slice4 = append(slice4, 15)
	slice4 = append(slice4, 16)
	fmt.Println("Slice 4:", slice4)
	fmt.Printf("Length: %d, Capacity: %d\n", len(slice4), cap(slice4))


	slice4 = append(slice4, 18)
	slice4 = append(slice4, 20)
	fmt.Println("Slice 4:", slice4)
	fmt.Printf("Length: %d, Capacity: %d\n", len(slice4), cap(slice4))


	// 4. From an Array
	array := [5]int{10, 11, 12, 13, 14}
	slice5 := array[1:4]
	fmt.Println("Slice 5:", slice5)
}