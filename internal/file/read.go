package file

import "io"

func ReadLines(r io.Reader) []string {
	lines := make([]string, 0)

	for line := range Lines(r) {
		lines = append(lines, line)
	}

	return lines
}
