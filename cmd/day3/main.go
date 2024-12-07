package main

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
)

const filename = "input/day3.txt"

func main() {
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		panic(err)
	}

	sum, sumWithEnabled := getMultipliesSum(file)

	fmt.Printf("Sum: %d, Sum with enabled: %d\n", sum, sumWithEnabled)
}

func findMultipliesByRegexp(s string, r *regexp.Regexp) int {
	matches := r.FindAllStringSubmatch(s, -1)

	res := 0
	enabled := true

	for _, match := range matches {
		if match[0] == "do()" {
			enabled = true
		} else if match[0] == "don't()" {
			enabled = false
		} else if enabled {
			first, _ := strconv.Atoi(match[2])
			second, _ := strconv.Atoi(match[3])
			res += first * second
		}
	}

	return res
}

func findMultiplies(s string) int {
	return findMultipliesByRegexp(s, regexp.MustCompile(`(mul\((\d+),(\d+)\))`))
}

func findMultipliesWithEnabled(s string) int {
	return findMultipliesByRegexp(s, regexp.MustCompile(`(mul\((\d+),(\d+)\))|(do\(\))|(don't\(\))`))
}

func getMultipliesSum(r io.Reader) (int, int) {
	text, err := io.ReadAll(r)
	if err != nil {
		panic(err)
	}

	return findMultiplies(string(text)), findMultipliesWithEnabled(string(text))
}
