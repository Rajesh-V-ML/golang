//JSON Unmarshall
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

type Employees struct {
	Employee []Employee `json:"users"`
}

type Employee struct {
	Name   string `json:"name"`
	Role   string `json:"role"`
	Age    int    `json:"age"`
	Social Social `json:"social"`
}

type Social struct {
	Mail string `json:"mail"`
	Num  string `json:"num"`
}

func main() {
	jsonFile, err := os.Open("emp.json")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Employee.json Opened")
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var Employees Employees
	json.Unmarshal(byteValue, &Employees)

	for i := 0; i < len(Employees.Employee); i++ {
		fmt.Println("Employee Name: " + Employees.Employee[i].Name)
		fmt.Println("Employee Role: " + Employees.Employee[i].Role)
		fmt.Println("Employee Age: " + strconv.Itoa(Employees.Employee[i].Age))
		fmt.Println("Employee Mail: " + Employees.Employee[i].Social.Mail)
		fmt.Println("Employee Num: " + Employees.Employee[i].Social.Num)

	}

}
