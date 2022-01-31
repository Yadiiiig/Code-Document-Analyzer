package main

import (
	"bytes"
	"fmt"
	"log"

	"github.com/Yadiiiig/Code-Document-Analyzer/internal"
	"github.com/Yadiiiig/Code-Document-Analyzer/searchers"
)

func main() {
	lines, err := searchers.Read("_example_code/main.go")
	if err != nil {
		log.Println(err)
	}

	//fmt.Println(lines)
	py := bytes.NewBufferString("")
	for _, v := range lines {
		internal.Splitter(&v)
		//fmt.Println(v.RawText)
		fmt.Fprintf(py, "['%s'], ", v.RawText)
		searchers.FindFunctions(&v)
	}

	fmt.Println(py.String())
}
