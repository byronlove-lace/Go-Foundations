package main

// Share memory by communicating - DJ Pike

import (
	"fmt"
)

func main() {
	// channels are the pipes between concurrent goroutines
	// They're like an internal FIFO queue
	// Some goroutines send vals to the queue and others receive from the queue
	// Goroutines transfer owenrship through this process tooj
	messages := make(chan string)

	// Here we create a channel and assign types
	binaryChannel := make(chan int, 10)
	close(binaryChannel)

	// add output of goroutine func into a channel with <-
	go func() { messages <- "ping" }()
	msg := <-messages

	fmt.Println(msg)
}
