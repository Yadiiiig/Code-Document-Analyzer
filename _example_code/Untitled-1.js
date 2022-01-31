{
    "name": "",
    "extension": "",
    "lines": [
        {
            "line": 0,
            "text": "package main",
            "length": 12,
            "functions": null,
            "variables": null
        },
        {
            "line": 1,
            "text": "",
            "length": 0,
            "functions": null,
            "variables": null
        },
        {
            "line": 2,
            "text": "import \"log\"",
            "length": 12,
            "functions": null,
            "variables": null
        },
        {
            "line": 3,
            "text": "",
            "length": 0,
            "functions": null,
            "variables": null
        },
        {
            "line": 4,
            "text": "// person is a struct type for a person records.",
            "length": 48,
            "functions": null,
            "variables": null
        },
        {
            "line": 5,
            "text": "type person struct {",
            "length": 20,
            "functions": null,
            "variables": null
        },
        {
            "line": 6,
            "text": "\tName   string",
            "length": 14,
            "functions": null,
            "variables": null
        },
        {
            "line": 7,
            "text": "\tPhone  string",
            "length": 14,
            "functions": null,
            "variables": null
        },
        {
            "line": 8,
            "text": "\tGender string",
            "length": 14,
            "functions": null,
            "variables": null
        },
        {
            "line": 9,
            "text": "}",
            "length": 1,
            "functions": null,
            "variables": null
        },
        {
            "line": 10,
            "text": "",
            "length": 0,
            "functions": null,
            "variables": null
        },
        {
            "line": 11,
            "text": "func main() {",
            "length": 13,
            "functions": null,
            "variables": null
        },
        {
            "line": 12,
            "text": "\tp := person{",
            "length": 13,
            "functions": null,
            "variables": null
        },
        {
            "line": 13,
            "text": "\t\tName:   \"Foo Bar\",",
            "length": 20,
            "functions": null,
            "variables": null
        },
        {
            "line": 14,
            "text": "\t\tPhone:  \"44864646\",",
            "length": 21,
            "functions": null,
            "variables": null
        },
        {
            "line": 15,
            "text": "\t\tGender: \"Male\",",
            "length": 17,
            "functions": null,
            "variables": null
        },
        {
            "line": 16,
            "text": "\t}",
            "length": 2,
            "functions": null,
            "variables": null
        },
        {
            "line": 17,
            "text": "\terr := register_person(p)",
            "length": 26,
            "functions": null,
            "variables": null
        },
        {
            "line": 18,
            "text": "\tif err != nil {",
            "length": 16,
            "functions": null,
            "variables": null
        },
        {
            "line": 19,
            "text": "\t\tlog.Println(err)",
            "length": 18,
            "functions": null,
            "variables": null
        },
        {
            "line": 20,
            "text": "\t}",
            "length": 2,
            "functions": null,
            "variables": null
        },
        {
            "line": 21,
            "text": "}",
            "length": 1,
            "functions": null,
            "variables": null
        },
        {
            "line": 22,
            "text": "",
            "length": 0,
            "functions": null,
            "variables": null
        },
        {
            "line": 23,
            "text": "// This is a function that adds a new person to a database.",
            "length": 59,
            "functions": null,
            "variables": null
        },
        {
            "line": 24,
            "text": "func register_person(p person) error {",
            "length": 38,
            "functions": null,
            "variables": null
        },
        {
            "line": 25,
            "text": "\t// pseudo code",
            "length": 15,
            "functions": null,
            "variables": null
        },
        {
            "line": 26,
            "text": "\tlog.Println(\"Added user to database\")",
            "length": 38,
            "functions": null,
            "variables": null
        },
        {
            "line": 27,
            "text": "\treturn nil",
            "length": 11,
            "functions": null,
            "variables": null
        },
        {
            "line": 28,
            "text": "}",
            "length": 1,
            "functions": null,
            "variables": null
        }
    ],
    "function_amount": 0,
    "variables_amount": 0,
    "markdown_version": ""
}