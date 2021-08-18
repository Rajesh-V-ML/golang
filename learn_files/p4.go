//call by value
//swap
//multiple value return
//input

package main

import "fmt"

func main() {
	fmt.Printf("Hello")

	var x, y int
	fmt.Println("Enter X & Y")
	fmt.Scanf("%d", &x)
	fmt.Scanf("%d", &y)

	fmt.Println(swap(x, y))

}

func swap(x, y int) (int, int) {
	temp := x
	x = y
	y = temp

	return x, y
}
