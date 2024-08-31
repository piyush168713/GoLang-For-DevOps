package main
import (
	"fmt"
)

func add(x, y int) int {
	return x+y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum*4/9;
	y = sum-x
	return
}

func sum_of_num(nums ...int) int {
	total := 0
	for _, num := range nums{
		total += num
	}
	return total
}

func main(){
	// Basic function
	result := add(3,4)
	fmt.Println("Result: ", result)

	// Multiple return values
	a, b := swap("hello", "world")
	fmt.Println("a, b = ", a, b)

	// Named return values
	sum := 100
	x, y := split(sum)
	fmt.Printf("The split of %d is: x = %d, y = %d", sum, x, y)

	// Variadic functions
	sum_result := sum_of_num(1,2,3,4)
	fmt.Println("\nThe sum is:", sum_result)
}