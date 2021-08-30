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

	printEmp(Emp1)
	printEmp(Emp2)
}

func printEmp(E Employee) {
	fmt.Println("Employee Code : ", E.ECode)
	fmt.Println("Employee Name : ", E.EName)
	fmt.Println("Employee Designation : ", E.EDesig)
	fmt.Println("Employee Salary : ", E.ESal)
	fmt.Println("Employee Score : ", E.EScore)
}
