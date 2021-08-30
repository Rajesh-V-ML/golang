//array input
package main

import "fmt"

func main() {
	fmt.Printf("Enter 5 Array Elements")
	var names [5]string
	for i := 0; i < 5; i++ {
		fmt.Printf("Enter %d th Element: ", i)
		fmt.Scanf("%s", &names[i])
		fmt.Println()
	}
	fmt.Println("Your Array is: ", names)

	// names[0] = "r"

	// fmt.Fscanf(])
	// names[1] = "a"
	// names[2] = "j"
	// names[3] = "r"
	// names[4] = "s"
	fmt.Println("array", names)
}
