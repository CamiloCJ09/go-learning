package main

import (
	"fmt"
)

func main() {
	var err error
	var salary int32 = 7000
	salaryError := SalaryError{}

	if salary <= 10000 {
		err = salaryError.Error(int(salary))
		if err != nil {
			fmt.Println(err)
		}
	}

}

type SalaryError struct{}

func (e *SalaryError) Error(value int) error {
	return fmt.Errorf("the salary entered does not reach the taxable minimum: %v ", value)
}
