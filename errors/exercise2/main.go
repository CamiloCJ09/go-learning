package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("Hello, playground")
	var salary int32 = 7000
	var err2 = &SalaryError{"The salary entered does not reach the taxable minimum"}

	var se error

	if salary <= 10000 {
		se = err2
	}

	fmt.Println(&se)
	fmt.Println(&err2)

	fmt.Println(errors.Is(se, err2))
	if errors.Is(se, err2) {
		fmt.Println(se)
	}

}

type SalaryError struct {
	msg string
}

func (e *SalaryError) Error() string {
	return e.msg
}
