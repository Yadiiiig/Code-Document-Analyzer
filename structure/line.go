package structure

type Line struct {
	Line      int        `json:"line"`     // Position of the line (vertical)
	RawText   string     `json:"raw_text"` // Complete string of line
	Text      []string   `json:"text"`
	Length    int        `json:"length"`    // Length of the text
	Functions []Function `json:"functions"` // Array of all the referenced functions
	Variables []Variable `json:"variables"` // Array of all the referenced variables
}

type Function struct {
	Name            string     `json:"name"`
	Comment         string     `json:"comment"`
	Variables       []Variable `json:"variables"`
	ReturnVariables []Variable `json:"return_variables"`
}

type Variable struct {
	Name       string      `json:"name"`
	Comment    string      `json:"comment"`
	References []Reference `json:"references"`
}

type Reference struct {
	Position Position `json:"position"`
}

type Position struct {
	Start int `json:"start"` // Start position of reference
	End   int `json:"end"`   // End position of reference
}
