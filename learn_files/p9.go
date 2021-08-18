//array input
package main

import "fmt"

func main() {
	fmt.Printf("Enter 5 Array Elements")
	var names [5]string
	for i := 0; i < 5; i++ {
		fmt.Printf("Enter %d th Element: ", i)
		fmt.Scanf("%s", &names[i])
	}
	fmt.Println("Your Array is: ", names)
}
