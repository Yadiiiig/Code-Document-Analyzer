package searchers

import (
	"github.com/Yadiiiig/Code-Document-Analyzer/structure"
)

func FindFunctions(l *structure.Line) ([]structure.Function, error) {
	for _, t := range l.Text {
		//fmt.Println(l.Text)
		if t == "func" {
			//fmt.Println(l.RawText)

		}
	}
	return nil, nil
}
