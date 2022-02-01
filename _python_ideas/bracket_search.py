# bracket_search.py

arr = [['package main'], [''], ['import "log"'], [''], ['// person is a struct type for a person records.'], 
    ['type person struct {'], ['  Name   string'], ['     Phone  string'], ['     Gender string'], ['}'], [''], 
    ['func main() {'], [' p := person{'], ['              Name:   "Foo Bar",'], ['           Phone:  "44864646",'], 
    ['                Gender: "Male",'], ['   }'], [' err := register_person(p)'], ['     if err != nil {'], 
    ['           log.Println(err)'], ['  }'], ['}'], [''], ['// This is a function that adds a new person to a database.'], 
    ['func register_person(p person) error {'], ['   // pseudo code'], ['        log.Println("Added user to database")'], 
    ['     return nil'], ['}']]

"""
Loop through each element in the array
Look for elements that start with func (starts definition of a function)
Stop....got some shit to do

Declare empty array/string/dict???
SAVE function name (could use string index [5:-2])
SAVE index of element that starts function definition
append current element (function declaration) to str/arr/dict
append every element following until we find a closing bracket***
    -use boolean to keep track of everything inside of brackets?
    -use bracket counter to count how many '{' are encountered   

"""


class func:
    defn=""
    index = ...
    def __init__(self, func_name):
        self.func_name = func_name

# main = func("main")

l=len(arr)
for i in range(l):
# for line in input:
    found_function = False
    if "".join(arr[i])[:4]=="func" and not found_function:
        found_function = True
        func_name = arr[i][5:-2]
        new_func = func_name(func_name)
        new_func.index = i

        if found_function:
            bracket_counter = 0
            new_func.defn += "(everything following the current element)"
            # func.defn+="\n".join(arr[i])

print(func.defn)








# some_function = "print(\"Hello world!\")"

# print(input)

# Outcome:
#
# Get the line height and start / end position for the function start and end
# From f to closing bracket


### Make a class for containing code contents, something to save function name into for example