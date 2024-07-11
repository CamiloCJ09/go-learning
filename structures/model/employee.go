package model

import "encoding/json"

type Employee struct {
	Id       int    `json:"id"`
	Position string `json:"position"`
	Person   Person `json:"person"`
}

func (e Employee) PrintEmployee() {
	employee, _ := json.MarshalIndent(e, "", "  ")
	println(string(employee))
}
