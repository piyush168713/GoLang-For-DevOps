package main
import "fmt"

func main(){
	// basic for loop
	// for i := 0; i < 5; i++ {
	// 	fmt.Println(i)
	// }

	// for loop as a while loop
	// i := 0
	// for i < 5 {
	// 	fmt.Println(i)
	// 	i++
	// }

	// infinite for loop
	// j := 0
	// for {
	// 	// some code
	// 	if(j > 100){
	// 		break;
	// 	}
	// 	fmt.Println(j)
	// 	j++
	// }

	// for loop with range (for slices, array, maps)
	// nums := []int{5,4,3,2,1}
	// for index, value := range nums {
	// 	fmt.Printf("Index: %d, Value: %d\n", index, value)
	// }

	// for loop with range (for strings)
	for index, value := range "Hello" {
		fmt.Printf("Index: %d, Value: %c\n", index, value)
	}
}