package main

import "testing"

func BenchmarkPoc(b *testing.B) {
	input := []string{
		"_example_code/main.go",
		"searchers/function.go",
	}
	for i := 0; i < b.N; i++ {
		poc(input)
	}
}
