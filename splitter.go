package ical

import (
	"bufio"
	"io"
	"strings"
)

func ContentLines(doc io.Reader) (out []string) {
	scanner := bufio.NewScanner(doc)
	line := ""
	for scanner.Scan() {
		token := scanner.Text()
		switch {
		case line == "":
			line = token
		case strings.HasPrefix(token, " "), strings.HasPrefix(token, "\t"):
			line += token[1:]
		default:
			out = append(out, line)
			line = token
		}
	}
	if line != "" {
		out = append(out, line)
	}
	return out
}
