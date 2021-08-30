//function call sample
package main

import "fmt"

func area(length, width float64) float64 {

	Ar := length * width
	return Ar
}

func main() {
	fmt.Printf("Area of rectangle is : %2.2f", area(15, 8))
}
