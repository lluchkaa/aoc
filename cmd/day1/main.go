package main

import (
	"fmt"
	"os"

	"github.com/lluchkaa/aoc/pkg/day1"
)

const filename = "input/day1.txt"

func main() {
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		panic(err)
	}

	first, second := day1.ReadLists(file)

	fmt.Printf("Difference: %d\n", day1.GetListsDifference(first, second))
	fmt.Printf("Similarity: %d\n", day1.GetSimilarityScore(first, second))
}
