package main

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/Yadiiiig/Code-Document-Analyzer/internal"
	"github.com/Yadiiiig/Code-Document-Analyzer/searchers"
	"github.com/Yadiiiig/Code-Document-Analyzer/structure"
)

func main() {
	input := []string{
		"_example_code/main.go",
		"searchers/function.go",
	}
	poc(input)
}

func poc(filenames []string) {
	for _, v := range filenames {
		lines, raw, err := searchers.Read(v)
		if err != nil {
			log.Println(err)
		}

		file := searchers.GetFileDetails(v, raw)
		var functions []structure.Function

		for k, v := range lines {
			internal.Splitter(&v)
			tmp_functions := searchers.FindFunctions(&v, k)
			if tmp_functions != nil {
				functions = append(functions, tmp_functions...)
			}
		}
		file.OriginalName = v
		file.Lines = lines
		file.Functions = functions
		file.FuncAmount = len(functions)

		result, _ := json.MarshalIndent(file, "", " ")

		_ = ioutil.WriteFile("__results/"+file.Name+".json", result, 0644)
	}
}
