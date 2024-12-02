package day1

import (
	"io"
	"sort"
	"strconv"
	"strings"

	"github.com/lluchkaa/aoc/internal/file"
	"github.com/lluchkaa/aoc/internal/slice"
)

func parseLine(line string) (int, int) {
	parts := strings.Split(line, "   ")

	first, _ := strconv.Atoi(parts[0])
	second, _ := strconv.Atoi(parts[1])

	return first, second
}

func ReadLists(r io.Reader) ([]int, []int) {
	first := make([]int, 0)
	second := make([]int, 0)

	for line := range file.Lines(r) {
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

func GetListsDifference(first []int, second []int) int {
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

func GetSimilarityScore(first []int, second []int) int {
	secondMap := slice.Frequency(second)

	return calculateSimilarityScore(first, secondMap)
}
