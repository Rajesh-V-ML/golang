//go routine
package main

import (
	"fmt"
	"time"
)

func hello() {
	fmt.Println("Hello world goroutine")
}

func main() {
	go hello()
	time.Sleep(1 * time.Second) // we wait till the function returns from hello
	fmt.Println("main function")
}
