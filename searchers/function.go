package searchers

import (
	"strings"

	"github.com/Yadiiiig/Code-Document-Analyzer/structure"
)

func FindFunctions(l *structure.Line, key int) []structure.Function {
	var res []structure.Function

	for _, t := range l.Text {
		if t == "func" {
			res = append(res, get_details(l.Text, l.RawText, key))
		}
	}
	if res != nil {
		return res
	} else {
		return nil
	}
}

func get_details(text []string, raw string, key int) structure.Function {
	var tmp []string
	var first_variable string
	var function_name string
	var function_variables []string
	var return_variables []string

	// Check if function has a variable
	if text[1][0:1] == "(" {
		tmp = strings.SplitN(raw, "(", 2)
		tmp = strings.SplitN(tmp[1], ")", 2)
		first_variable = tmp[0]

		tmp = strings.SplitN(tmp[1], "(", 2)
		function_name = tmp[0]
	} else {
		tmp = strings.SplitN(raw, " ", 2)
		tmp = strings.SplitN(tmp[1], "(", 2)
		function_name = tmp[0]
	}

	// Check if function has function parameters
	if tmp[1][0:1] != ")" {
		tmp = strings.SplitN(tmp[1], ")", 2)
		function_variables = strings.Split(tmp[0], ", ")
	}

	tmp = strings.SplitN(tmp[1], "(", 2)

	if strings.Contains(tmp[0], "{") {

	} else if strings.Contains(tmp[1], ",") {
		tmp = strings.SplitN(tmp[1], ")", 2)
		return_variables = strings.Split(tmp[0], ", ")
	} else if !strings.Contains(tmp[0], ")") {
		tmp = strings.Split(tmp[0], " ")
		return_variables = []string{tmp[1]}
	}

	var vars []structure.Variable
	if len(function_variables) != 0 {
		for _, v := range function_variables {
			tmp_o := strings.Split(v, " ")
			vars = append(vars, structure.Variable{
				Name: tmp_o[0],
				Type: tmp_o[1],
			})
		}
	}

	var r_vars []structure.Variable
	if len(return_variables) != 0 {
		for _, v := range return_variables {
			vars = append(vars, structure.Variable{
				Type: v,
			})
		}
	}

	return structure.Function{
		Name:            function_name,
		OOPVariable:     first_variable,
		Variables:       vars,
		ReturnVariables: r_vars,
		LineHeight:      key + 1,
	}
}
