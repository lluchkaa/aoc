package main

import (
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/lluchkaa/aoc/internal/reader"
	"github.com/lluchkaa/aoc/internal/slice"
)

const filename = "input/day1.txt"

func main() {
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		panic(err)
	}

	first, second := readLists(file)

	fmt.Printf("Difference: %d\n", getListsDifference(first, second))
	fmt.Printf("Similarity: %d\n", getSimilarityScore(first, second))
}

func parseLine(line string) (int, int) {
	parts := strings.Split(line, "   ")

	first, _ := strconv.Atoi(parts[0])
	second, _ := strconv.Atoi(parts[1])

	return first, second
}

func readLists(r io.Reader) ([]int, []int) {
	first := make([]int, 0)
	second := make([]int, 0)

	for line := range reader.Lines(r) {
		x, y := parseLine(line)

		first = append(first, x)
		second = append(second, y)
	}

	return first, second
}

func calculateDifference(first []int, second []int) int {
	res := 0

	for i, x := range first {
		if x < second[i] {
			res += second[i] - x
		} else {
			res += x - second[i]
		}

	}

	return res
}

func getListsDifference(first []int, second []int) int {
	sort.Ints(first)
	sort.Ints(second)

	return calculateDifference(first, second)
}

func calculateSimilarityScore(first []int, secondMap map[int]int) int {
	res := 0

	for _, x := range first {
		res += secondMap[x] * x
	}

	return res
}

func getSimilarityScore(first []int, second []int) int {
	secondMap := slice.Frequency(second)

	return calculateSimilarityScore(first, secondMap)
}
