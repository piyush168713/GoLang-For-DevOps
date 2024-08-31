package main

import (
	"fmt"
	"time"
)

func printCount(label string, count int, ch chan int) {
	for i := 0; i < count; i++ {
		ch <- i
		fmt.Println(label, i)
		time.Sleep(time.Millisecond * 500)
	}
	close(ch)
}

func main(){
	ch := make(chan int)
	go printCount("goroutine", 5, ch)

	for value := range ch {
		fmt.Println("Main recieved:", value)
	}
}