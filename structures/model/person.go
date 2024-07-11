package model

type Person struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	DateOfBirth string `json:"date_of_birth"`
}

func (p Person) PrintPerson() {
	println(p.Id)
	println(p.Name)
	println(p.DateOfBirth)
}
