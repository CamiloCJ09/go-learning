package main

import (
	"errors"
	"fmt"
)

const (
	minHours  int = 80
	limSalary int = 150000
)

func main() {
	workedHours := 125
	hoursPrice := 10

	salary, err := CalculateSalary(workedHours, float64(hoursPrice))

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(salary)

}

func CalculateSalary(workedHours int, hoursPrice float64) (float64, error) {

	var calculatedSalary float64 = float64(workedHours) * hoursPrice

	if salaryOver150k(calculatedSalary) {
		calculatedSalary = calculatedSalary * 0.9
	}

	if hoursLowerThan80(workedHours) {
		return 0, errors.New("the worker cannot have worked less than 80 hours per month")
	}

	return calculatedSalary, nil
}

func salaryOver150k(salary float64) bool {
	return salary > float64(limSalary)
}

func hoursLowerThan80(workedHours int) bool {
	return workedHours < minHours
}
