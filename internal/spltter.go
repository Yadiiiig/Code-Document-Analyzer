package internal

import (
	"strings"

	"github.com/Yadiiiig/Code-Document-Analyzer/structure"
)

func Splitter(line *structure.Line) {
	tmp := strings.Split(line.RawText, " ")
	line.Text = tmp
}
