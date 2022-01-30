package main

import "log"

// cda person is a struct type for a person records.
type person struct {
	Name   string
	Phone  string
	Gender string
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
}

/*
cda
This is a function that adds a new person to a database.
Argument:
- person
Returns:
- error
*/
func register_person(p person) error {
	// pseudo code
	log.Println("Added user to database")
	return nil
}
