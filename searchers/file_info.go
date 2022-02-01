package searchers

import (
	"bytes"
	"strings"

	"github.com/Yadiiiig/Code-Document-Analyzer/structure"
)

func GetFileDetails(filename string, raw []string) structure.File {
	var file structure.File
	var b bytes.Buffer
	var line_height int

	path := strings.Split(filename, "/")
	tmp_name := strings.Split(path[len(path)-1], ".")

	for _, v := range raw {
		b.WriteString(v)
		line_height += 1
	}

	file.Name = path[len(path)-1]
	file.Extension = tmp_name[1]
	file.Raw = b.String()
	file.LineHeight = line_height

	return file
}
