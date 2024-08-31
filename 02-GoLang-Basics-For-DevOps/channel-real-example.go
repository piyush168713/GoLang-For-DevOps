package main

import (
	"fmt"
	"time"
)

type Order struct{
	ID int
	Status string
}

func processOrder(orderId int, statusChannel chan Order){
	time.Sleep(time.Second*2)
	// Sending processed order status back through the channel
	statusChannel <- Order{ID: orderId, Status: "Completed"}
}

func main(){
	// creating a channel for communicating order statuses
	statusChannel := make(chan Order)

	// Starting goroutines for processing orders, i.e. it creates 6 goroutines
	for i := 0; i <= 5; i++ {
		go processOrder(i, statusChannel)  
	}

	// Recieving and Printing order Statuses
	for i := 0; i < 5; i++ {
		processOrder := <-statusChannel;
		fmt.Printf("Order %d processed: %s\n", processOrder.ID, processOrder.Status)
	}
}