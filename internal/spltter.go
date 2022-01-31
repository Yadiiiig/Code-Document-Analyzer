package internal

import (
	"strings"

	"github.com/Yadiiiig/Code-Document-Analyzer/structure"
)

func Splitter(line *structure.Line) {
	//fmt.Println(line.RawText)
	tmp := strings.Split(line.RawText, " ")
	line.Text = tmp

}
