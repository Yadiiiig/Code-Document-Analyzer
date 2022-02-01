package structure

type Function struct {
	Name            string     `json:"name"`
	Comment         string     `json:"comment"`
	OOPVariable     string     `json:"oop_variable", omitempty`
	Variables       []Variable `json:"variables"`
	ReturnVariables []Variable `json:"return_variables"`
	LineHeight      int        `json:""line_height`
}
