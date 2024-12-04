package main

import (
	"fmt"
	"os"

	"github.com/lluchkaa/aoc/internal/file"
	"github.com/lluchkaa/aoc/pkg/day4"
)

const filename = "input/day4.txt"

func main() {
	inputFile, err := os.Open(filename)
	defer inputFile.Close()
	if err != nil {
		panic(err)
	}

	lines := file.ReadLines(inputFile)
	count := day4.GetCountOfOccurrences(lines, "XMAS")
	xCount := day4.GetCountOfXMases(lines)

	fmt.Printf("Count: %d; X Count: %d\n", count, xCount)
}
