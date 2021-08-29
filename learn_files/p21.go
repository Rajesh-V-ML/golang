//JSON Marshal
package main

import (
	"encoding/json"
	"io/ioutil"
)

type Salary struct {
	Basic, Allowance, PF float64
}

type Employee struct {
	FName, LName, Mail string
	Age                int
	MonthlySal         []Salary
}

func main() {
	data := Employee{
		FName: "Rajesh",
		LName: "V",
		Mail:  "rajesh.v@maplelabs.com",
		Age:   24,
		MonthlySal: []Salary{
			Salary{
				Basic:     20000,
				Allowance: 1000,
				PF:        2000,
			},
			Salary{
				Basic:     25000,
				Allowance: 1500,
				PF:        2200,
			},
		},
	}
	file, _ := json.MarshalIndent(data, "", " ")
	_ = ioutil.WriteFile("employee.json", file, 0644)
}
