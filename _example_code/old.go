package main

import "log"

// person is a struct type for a person records.
type person struct {
	Name     string
	Phone    string
	Gender   string
	Birthday string
}

type hello {
	World string
}

func main() {
	p := person{
		Name:   "Foo Bar",
		Phone:  "44864646",
		Gender: "Male",
	}
	err := register_person(p)
	if err != nil {
		log.Println(err)
	}
	p.add_birthday("1998")
}

// This is a function that adds a new person to a database.
func register_person(p person) error {
	// pseudo code
	log.Println("Added user to database")
	return nil
}

func (person) add_birthday(date string) {
	person.Birthday = date
}

func (person) foo(date string, age int) error {
	return nil
}

func (hello) foo(date string) (string, error) {
	return "", nil
}
