package structure

type File struct {
	Name            string `json:"name"`
	Extension       string `json:"extension"`
	Lines           []Line `json:"lines"`
	FuncAmount      int    `json:"function_amount"`
	VarAmount       int    `json:"variables_amount"`
	MarkdownVersion string `json:"markdown_version"`
}
