//anonymous function

package main

import (
	f "fmt"
)

func getSequence() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

func main() {
	nextNumber := getSequence()

	f.Println(nextNumber())
	f.Println(nextNumber())
	f.Println(nextNumber())
	f.Println(nextNumber())

	nextNumber1 := getSequence()
	f.Println(nextNumber1())
	f.Println(nextNumber1())
}
