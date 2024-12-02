package day2

import (
	"io"

	"github.com/lluchkaa/aoc/internal/file"
)

func GetCountOfSafeReports(r io.Reader) int {
	count := 0
	for line := range file.Lines(r) {
		if isSafe(parseLine(line)) {
			count += 1
		}
	}

	return count
}

func GetCountOfSafeReportsWithDampener(r io.Reader) int {
	count := 0
	for line := range file.Lines(r) {
		if isSafeWithDampener(parseLine(line)) {
			count += 1
		}
	}

	return count
}
