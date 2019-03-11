package main

import (
	"fmt"
	"math/rand"
	"time"
)

// producer
func producer(header string, channel chan<- string) {
	// infinite for creating random data
	for {
		// send back to channel
		channel <- fmt.Sprintf("%s: %v", header, rand.Int31())

		// sleep for 1 second
		time.Sleep(time.Second)
	}
}

// consumer
func consumer(channel <-chan string) {
	// infinite for retrieving data
	for {
		// waiting for channel response
		message := <-channel

		// print it out
		fmt.Println(message)
	}
}

func main() {
	// create a string typed channel
	channel := make(chan string)

	// create goroutine for producer
	go producer("cat", channel)
	go producer("dog", channel)

	// create consumer
	consumer(channel)
}
