package main

import (
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/lluchkaa/aoc/internal/reader"
)

const filename = "input/day7.txt"

func main() {
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		panic(err)
	}

	equations := readData(file)

	part1, part2 := Solve(equations)
	fmt.Printf("Part 1: %d\nPart 2: %d\n", part1, part2)
}

type equation struct {
	result int
	data   []int
}

func digits(num int) int {
	res := int(math.Log10(float64(num))) + 1
	if l := len(strconv.Itoa(num)); l != res {
		panic(fmt.Sprintf("res = %d, len = %d, num = %d", res, l, num))
	}
	return res
}

func (e equation) tryToSolve(left int, i int, concatEnabled bool) bool {
	if i >= len(e.data) {
		return false
	}

	if i == 0 {
		return left == e.data[i]
	}

	if left < 0 {
		return false
	}

	power := int(math.Pow10(digits(e.data[i])))

	return e.tryToSolve(left-e.data[i], i-1, concatEnabled) ||
		left%e.data[i] == 0 && e.tryToSolve(left/e.data[i], i-1, concatEnabled) ||
		concatEnabled && e.data[i] == left%power && e.tryToSolve(left/power, i-1, concatEnabled)
}

func (e equation) canBeSolved(concatEnabled bool) bool {
	return e.tryToSolve(e.result, len(e.data)-1, concatEnabled)
}

func readData(r io.Reader) []equation {
	equations := make([]equation, 0)

	for line := range reader.Lines(r) {
		e := equation{data: []int{}}
		colonIndex := strings.IndexByte(line, ':')

		e.result, _ = strconv.Atoi(line[:colonIndex])

		parts := strings.Split(line[colonIndex+2:], " ")

		for _, part := range parts {
			num, _ := strconv.Atoi(part)
			e.data = append(e.data, num)
		}

		equations = append(equations, e)
	}

	return equations
}

func Solve(equations []equation) (int, int) {
	part1 := 0
	part2 := 0

	for _, e := range equations {
		if e.canBeSolved(false) {
			part1 += e.result
			part2 += e.result
		} else if e.canBeSolved(true) {
			part2 += e.result
		}
	}

	return part1, part2
}
