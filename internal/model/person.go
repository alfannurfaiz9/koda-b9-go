package model

import "fmt"

type Person struct {
	Name    string
	Address string
	Phone   string
}

func NewPerson(name, address, phone string) *Person {
	return &Person{
		Name:    name,
		Address: address,
		Phone:   phone,
	}
}

func (person *Person) GetPerson() {
	fmt.Printf("Name = %s \nAddress = %s \nPhone = %s \n", person.Name, person.Address, person.Phone)
}

func (person *Person) Greet() {
	fmt.Println("Hello,", person.Name)
}

func (person *Person) SetName(newName string) {
	person.Name = newName
}
