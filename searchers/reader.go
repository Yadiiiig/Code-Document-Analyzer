package searchers

import (
	"bufio"
	"os"

	"github.com/Yadiiiig/Code-Document-Analyzer/structure"
)

func Read(fn string) ([]structure.Line, []string, error) {
	file, err := os.Open(fn)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	count := -1
	var lines []structure.Line
	var raw []string
	for scanner.Scan() {
		count += 1
		text := scanner.Text()
		line := structure.Line{
			Line:    count,
			RawText: text,
			Length:  len(text),
		}

		lines = append(lines, line)
		raw = append(raw, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, err
	}

	return lines, raw, nil
}
