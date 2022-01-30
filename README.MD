# Code Document Analyzer

- Scan over every file
- Starts reading the comment once a symbol is typed (this symbol can be referenced)
- Reads the summary of a function, arguments -, return values explanations and stores it in a public database
- Documented code is publicly available on the server (*)
- Backend groups/bind everything together in markdown document (**)
- Website/API calls will be able to be used to retrieve the documentation (***)

(* Can also be privatly hosted, ...)

(** At the beginning only comments will get analyzed, together with some basic 
clickthrough to variables)

(*** In the future it can be linked to a vsc extension for example.
If they link their repository, the backend automaticly commits it to their repo,
as a markdown file)


# Documentation

 **Overview:**
 `idea/main.go`
<hr></hr>

# **[RegisterPerson](RegisterPerson)**

```
register_person()
-----------------
This is a function that adds a new person to a database.
-----------------
Found usecases:

p  := person{
	Name:  "Foo Bar",
	Phone:  "44864646",
	Gender:  "Male",
}
err  :=  register_person(p)
if err !=  nil  {
	log.Println(err)
}

```
**Variables**
[Person (type person)](#Person)

**Returning Variables**
[Error](error.readme)

<hr></hr>

# **[Person](Person)**
```
person struct
-----------------
person is a struct type for a person records.
-----------------
Found usecases:

p  := person{
	Name:  "Foo Bar",
	Phone:  "44864646",
	Gender:  "Male",
}
```
**Variables**
[Name (string)](Name.readme)
[Phone (string)](Phone.readme)
[Gender (string)](Gender.readme)
