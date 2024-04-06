package main

import (
	"fmt"
	"time"
)

func listenToChan(ch chan int) {
	for {
		i := <-ch

		fmt.Printf("Got %d from channel!\n", i)

		//pretend to work on some task for a second
		time.Sleep(1 * time.Second)
	}
}

func main() {
	ch := make(chan int, 10) //buffer of 10. that makes it a buffered channel

	go listenToChan(ch)

	for i := 0; i < 100; i++ {
		fmt.Printf("Sending %d to channel!\n", i)
		ch <- i
		fmt.Printf("Sent %d to channel!\n", i)
	}
}
