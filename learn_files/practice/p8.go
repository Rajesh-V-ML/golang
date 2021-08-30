//array input
package main

import "fmt"

func main() {
	fmt.Printf("Enter the Array Size")
	var size int
	fmt.Scanln(&size)
	var names = make([]string, size)
	for i := 0; i < size; i++ {
		fmt.Printf("Enter %d th Element: ", i)
		fmt.Scanf("%s", &names[i])
	}
	fmt.Println("Your Array is: ", names)
}
