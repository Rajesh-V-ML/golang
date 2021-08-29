package main

import "fmt"

func main() {
	message := make(chan string)

	go func() { message <- "Hello" }() //anonymous function to send to channel
	//	go test(message)
	msg := <-message //recieve from channel
	fmt.Println(msg)
}
