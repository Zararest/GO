package main

import (
	"fmt"
	"temperature/pkg/interval"
)

func checkInputErr(err error) {
	if err != nil {
		panic("invalid input")
	}
}

func parseDepartment() {
	var numOfEmployees uint
	_, err := fmt.Scanln(&numOfEmployees)
	checkInputErr(err)

	interv := interval.Interval{From: 15, To: 30}
	for i := uint(0); i < numOfEmployees; i++ {
		var constraint string
		_, err = fmt.Scan(&constraint)
		checkInputErr(err)

		var t uint
		_, err = fmt.Scanln(&t)
		checkInputErr(err)

		err = interv.AddConstraint(constraint, t)
		checkInputErr(err)

		fmt.Println(interv.GetOptimal())
	}
}

func main() {
	var numOfDepartments uint
	_, err := fmt.Scanln(&numOfDepartments)
	checkInputErr(err)

	for i := uint(0); i < numOfDepartments; i++ {
		parseDepartment()
	}
}
