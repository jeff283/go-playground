package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("-- CHANNELS --")
	fmt.Println()

	workingChannel()
	fmt.Println()

	loopChannel()
	fmt.Println()

	bufferedChannels()
	fmt.Println()

}

func workingChannel() {
	var c = make(chan int)
	go process(c)
	fmt.Println("Channel Value: ", <-c)

}

func process(c chan int) {
	c <- 1

}

func brokenChannel() {
	// Make a channel
	// Holds values, is thread safe (go routines), listen for data
	// This channel is unbuffered i.e only one value in it at a time
	var c = make(chan int)

	// Add value to channel
	c <- 1

	// Get value from channel
	var v = <-c

	// Will error out
	// This is beacuse channels unbuffered channels need to give the value of to somewhere
	// we assign one and we wait for someone to read, but since our read should execute after that, we get into a deadlock at `c <- 1`
	fmt.Println("Channel Value: ", v)
}

func loopChannel() {
	fmt.Println("Looping Channel: ")
	var c = make(chan int)
	// Start the sender
	go loopProcess(c)

	// Receive them one by one as they come
	// Not doing so ends up in a deadlock
	for i := range c {
		fmt.Println(i)
	}

}
func loopProcess(c chan int) {
	// Say no more values are coming
	// Defer to tell the program to do this just before returning the function
	defer close(c)
	// Assing continous values to the loop
	for i := range 5 {
		c <- i
	}

	// // Close the channel for the last value so that we can return
	// close(c)
}

func bufferedChannels() {
	fmt.Println("Buffered Channels:")

	// Create a buffered channels with 5 slots
	// Lets the bufferProcess accumulate values without needing them to be read immediately
	// More like giving the channel a bucket for putting its things before they are delivered
	var c = make(chan int, 5)
	go bufferProcess(c)

	for i := range c {
		fmt.Println(i)
		time.Sleep(time.Second * 1)
	}
}

func bufferProcess(c chan int) {
	defer close(c)
	for i := range 5 {
		c <- i
	}
	fmt.Println("Exiting Buffer Process")

}
