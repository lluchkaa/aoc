package reader

import (
	"io"
)

func ReadLines(r io.Reader) []string {
	lines := make([]string, 0)

	for line := range Lines(r) {
		lines = append(lines, line)
	}

	return lines
}

func Dimensions(r io.Reader) (int, int) {
	rows := 0
	cols := 0

	for line := range Lines(r) {
		rows++
		if cols < len(line) {
			cols = len(line)
		}
	}

	return cols, rows
}
