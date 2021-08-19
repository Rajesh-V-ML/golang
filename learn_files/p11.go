//Structure
package main

import "fmt"

type Employee struct {
	ECode  int
	EName  string
	EDesig string
	ESal   int
	EScore float32
}

func main() {

	//Declare Variables
	var Emp1 Employee
	var Emp2 Employee

	//Emp1 Details
	Emp1.ECode = 10217
	Emp1.EName = "Rajesh"
	Emp1.EDesig = "Associate Software Engineer"
	Emp1.ESal = 35000
	Emp1.EScore = 4.5

	//Emp2 Details
	Emp2.ECode = 10218
	Emp2.EName = "Sam"
	Emp2.EDesig = "Senior Software Engineer"
	Emp2.ESal = 50000
	Emp2.EScore = 4.8

	fmt.Printf("Employee 1 Details: \n Code: %d Name: %s Designation: %s Salary: %d Score: %0.2f", Emp1.ECode, Emp1.EName, Emp1.EDesig, Emp1.ESal, Emp1.EScore)
	fmt.Printf("\nEmployee 2 Details: \n Code: %d Name: %s Designation: %s Salary: %d Score: %0.2f", Emp2.ECode, Emp2.EName, Emp2.EDesig, Emp2.ESal, Emp2.EScore)

}
