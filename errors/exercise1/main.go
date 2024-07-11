package main

import "fmt"

func main() {
	fmt.Println("Hello, playground")
	var salary int32 = 700000

	if err := validateSalary(salary); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Must pay taxes")

}

type SalaryError struct {
	Msg string
}

func (e *SalaryError) Error() string {
	return e.Msg
}

func validateSalary(salary int32) error {
	if salary < 150000 {
		return &SalaryError{"The salary entered does not reach the taxable minimum"}
	}
	return nil
}
