package structure

type File struct {
	Name            string     `json:"name"`
	OriginalName    string     `json:"original_name"`
	Extension       string     `json:"extension"`
	LineHeight      int        `json:"amount_lines"`
	FuncAmount      int        `json:"function_amount"`
	VarAmount       int        `json:"variables_amount"`
	Functions       []Function `json:"functions"`
	Lines           []Line     `json:"lines"`
	Raw             string     `json:"raw_bytes"`
	MarkdownVersion string     `json:"markdown_version"`
}
