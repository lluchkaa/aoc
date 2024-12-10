package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/lluchkaa/aoc/internal/reader"
)

const filename = "input/day6.txt"

func main() {
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		panic(err)
	}

	start, lines := readData(file)

	visitedCount, visited := Part1(start, lines)

	fmt.Printf("Part 1: %d\n", visitedCount)
	fmt.Printf("Part 2: %d\n", Part2(start, lines, visited))
}

type position = [2]int

type direction = [2]int

var startDirection direction = direction{0, -1}

type empty = struct{}

func turnRight(d direction) direction {
	return direction{-d[1], d[0]}
}

func isOutOfBounds(p position, h, w int) bool {
	return p[0] < 0 || p[1] < 0 || p[0] >= w || p[1] >= h
}

func readData(r io.Reader) (position, []string) {
	start := position{-1, -1}

	lines := make([]string, 0)

	for i, line := range reader.Lines2(r) {
		if index := strings.IndexByte(line, '^'); index >= 0 {
			start = position{index, i}
		}
		lines = append(lines, line)
	}

	return start, lines
}

func makeStep(p position, d direction, lines []string) (position, direction, bool) {
	w, h := len(lines[0]), len(lines)

	for i := 0; i < 4; i++ {
		next := position{p[0] + d[0], p[1] + d[1]}

		if isOutOfBounds(next, h, w) {
			return next, d, false
		}

		if lines[next[1]][next[0]] != '#' {
			return next, d, true
		}
		d = turnRight(d)
	}

	panic("cannot move")
}

func travel(p position, d direction, lines []string) (map[position]map[direction]empty, bool, bool) {
	visited := make(map[position]map[direction]empty)
	visited[p] = map[direction]empty{d: {}}

	inBounds := true
	inLoop := false

	for {
		p, d, inBounds = makeStep(p, d, lines)
		if !inBounds {
			break
		} else if _, ok := visited[p][d]; ok {
			inLoop = true
			break
		}

		if visited[p] == nil {
			visited[p] = map[direction]empty{}
		}
		visited[p][d] = empty{}
	}

	return visited, inBounds, inLoop
}

func Part1(p position, lines []string) (int, map[position]map[direction]empty) {
	visited, _, _ := travel(p, startDirection, lines)
	return len(visited), visited
}

func Part2(startPos position, lines []string, visited map[position]map[direction]empty) int {
	res := 0
	for pos := range visited {
		if pos == startPos {
			continue
		}

		exchangedLine := []byte(lines[pos[1]])
		exchangedLine[pos[0]] = '#'
		lines[pos[1]] = string(exchangedLine)

		_, _, isLoop := travel(startPos, startDirection, lines)
		if isLoop {
			res++
		}

		exchangedLine[pos[0]] = ' '
		lines[pos[1]] = string(exchangedLine)
	}

	return res
}
