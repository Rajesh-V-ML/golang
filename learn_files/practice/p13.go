//range & map

package main

import "fmt"

func main() {
	//range
	number := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}

	for i := range number {
		fmt.Printf("%d\t", number[i])
	}

	fmt.Printf("\nMap Function \n")
	//map
	nameage := map[string]int{"Rajesh": 24, "Ritika": 22}
	for name := range nameage {
		fmt.Println(name, nameage[name])
	}
}
