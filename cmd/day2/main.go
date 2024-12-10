package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/lluchkaa/aoc/internal/reader"
)

const filename = "input/day2.txt"

func main() {
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Safe reports count: %d\n", getCountOfSafeReports(file))
	file.Seek(0, 0)

	fmt.Printf("Safe reports with Dampener count: %d\n", getCountOfSafeReportsWithDampener(file))
}

type report = []int

func parseLine(s string) report {
	parts := strings.Split(s, " ")
	report := make(report, len(parts))

	for i, x := range parts {
		report[i], _ = strconv.Atoi(x)
	}

	return report
}

func isAscSafe(x, y int) bool {
	return y > x && y-x <= 3
}

func isDescSafe(x, y int) bool {
	return isAscSafe(y, x)
}

func isSafeFunc(r report, comparator func(x, y int) bool) bool {
	if len(r) < 2 {
		return true
	}

	for i := 1; i < len(r); i++ {
		if !comparator(r[i-1], r[i]) {
			return false
		}
	}

	return true
}

func isSafe(r report) bool {
	if len(r) < 2 {
		return true
	}

	comparator := isAscSafe
	if r[0] > r[1] {
		comparator = isDescSafe
	}

	return isSafeFunc(r, comparator)
}

func isSafeWithDampenerFunc(r report, comparator func(x, y int) bool) bool {
	if len(r) < 2 {
		return true
	}

	for i := 1; i < len(r); i++ {
		if !comparator(r[i-1], r[i]) {
			if isSafeFunc(append([]int{r[i-1]}, r[i+1:]...), comparator) {
				return true
			}
			withoutPrevious := []int{}
			if i >= 2 {
				withoutPrevious = append(withoutPrevious, r[i-2:i-1]...)
			}
			withoutPrevious = append(withoutPrevious, r[i:]...)
			return isSafeFunc(withoutPrevious, comparator)
		}
	}

	return true
}

func isSafeWithDampenerBruteForceFunc(r report, comparator func(x, y int) bool) bool {
	arr := make([]int, len(r)-1)

	for i := 0; i < len(r); i++ {
		copy(arr[:i], r[:i])
		copy(arr[i:], r[i+1:])
		if isSafeFunc(arr, comparator) {
			return true
		}
	}

	return false
}

func isSafeWithDampener(r report) bool {
	return isSafe(r) || isSafeWithDampenerFunc(r, isAscSafe) || isSafeWithDampenerFunc(r, isDescSafe)
}

func getCountOfSafeReports(r io.Reader) int {
	count := 0
	for line := range reader.Lines(r) {
		if isSafe(parseLine(line)) {
			count += 1
		}
	}

	return count
}

func getCountOfSafeReportsWithDampener(r io.Reader) int {
	count := 0
	for line := range reader.Lines(r) {
		if isSafeWithDampener(parseLine(line)) {
			count += 1
		}
	}

	return count
}
