//array
package main

import "fmt"

func main() {
	var names [10]int
	for i := 0; i < 10; i++ {
		names[i] = i + 1
	}

	for j := 0; j < 10; j++ {
		fmt.Printf("Elements[%d] = %d\n", j, names[j])
	}

}
