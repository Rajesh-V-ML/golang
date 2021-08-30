//slicing
package main

import f "fmt"

func main() {
	//slice
	var numbers = make([]int, 10, 10)
	for i := 0; i < cap(numbers); i++ {
		numbers[i] = i
	}

	f.Println(numbers)

	f.Println("Append ")
	numbers = append(numbers, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	f.Println(numbers)
}
