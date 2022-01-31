# bracket_search.py

input = ['package main'], [''], ['import "log"'], [''], ['// person is a struct type for a person records.'], 
    ['type person struct {'], ['  Name   string'], ['     Phone  string'], ['     Gender string'], ['}'], [''], 
    ['func main() {'], [' p := person{'], ['              Name:   "Foo Bar",'], ['           Phone:  "44864646",'], 
    ['                Gender: "Male",'], ['   }'], [' err := register_person(p)'], ['     if err != nil {'], 
    ['           log.Println(err)'], ['  }'], ['}'], [''], ['// This is a function that adds a new person to a database.'], 
    ['func register_person(p person) error {'], ['   // pseudo code'], ['        log.Println("Added user to database")'], 
    ['     return nil'], ['}']

print(input)

# Outcome:
#
# Get the line height and start / end position for the function start and end