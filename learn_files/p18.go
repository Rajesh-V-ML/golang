package main

import (
	"fmt"
)

func hello(done chan bool) {
	fmt.Println("HELLO FROM GO ROUTINE")
	done <- true //send to channel
}

func receive(receive bool, done2 chan bool) {
	if receive == true {
		fmt.Println("FROM RECEIVE")
		done2 <- true
	}
}

func main() {

	done := make(chan bool)
	done2 := make(chan bool)

	go hello(done)
	tosend := <-done // receive from channel
	go receive(tosend, done2)
	<-done2
	fmt.Println("HELLO FROM MAIN")
}

//https://medium.com/wesionary-team/understanding-go-routine-and-channel-b09d7d60e575
