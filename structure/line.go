package structure

type Line struct {
	Line      int        `json:"line"`     // Position of the line (vertical)
	RawText   string     `json:"raw_text"` // Complete string of line
	Text      []string   `json:"text"`
	Length    int        `json:"length"`    // Length of the text
	Variables []Variable `json:"variables"` // Array of all the referenced variables
}

type Variable struct {
	Name       string      `json:"name"`
	Type       string      `json:"type"`
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
