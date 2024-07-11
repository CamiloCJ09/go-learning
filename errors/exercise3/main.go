package main

import (
	"errors"
	"fmt"
)

func main() {
	var salary int32 = 7000
	salaryError := errors.New("the salary entered does not reach the taxable minimum")
	var err error

	if salary <= 10000 {
		err = salaryError
	}
	if errors.Is(err, salaryError) {
		fmt.Println(err)
	}

}

type SalaryError struct {
}
