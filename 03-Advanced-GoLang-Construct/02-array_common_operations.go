package main

import(
	"fmt"
)

func modifyArr(arr [5]int, index int, value int) [5]int{
	arr[index] = value
	return arr
}

func main(){
	arr := [5]int{1,2,3,4,5}

	fmt.Println("The length of the array is:", len(arr))

	var copiedArr = arr;
	fmt.Println("Copied Array:", copiedArr)

	modifiedArr := modifyArr(arr, 2, 10)
	fmt.Println("Modified Array:", modifiedArr)

	isEqual := arr == modifiedArr
	fmt.Println("Are original and modified arrays equal?", isEqual)

	isCopiedEqual := arr == copiedArr
	fmt.Println("Are original and copied arrays equal?", isCopiedEqual)

}