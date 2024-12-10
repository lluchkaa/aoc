package main

import (
	"fmt"
	"os"

	"github.com/lluchkaa/aoc/internal/ds/queue"
	"github.com/lluchkaa/aoc/internal/reader"
)

const filename = "input/day10.txt"

func main() {
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		panic(err)
	}

	lines := reader.ReadLines(file)

	fmt.Printf("Part 1: %d\n", part1(lines))
	fmt.Printf("Part 2: %d\n", part2(lines))
}

type empty = struct{}
type position = [2]int

var directions = [...][2]int{
	{-1, 0},
	{0, 1},
	{1, 0},
	{0, -1},
}

type cell struct {
	val byte
	x   int
	y   int
}

func walk(lines []string, complete func(position, position)) {
	m, n := len(lines), len(lines[0])

	for i := range lines {
		for j := range lines[i] {
			if lines[i][j] != '0' {
				continue
			}

			q := queue.New[cell]()
			q.Push(cell{lines[i][j], j, i})

			for c, ok := q.Pop(); ok; c, ok = q.Pop() {
				for _, d := range directions {
					x, y := c.x+d[1], c.y+d[0]

					if x < 0 || y < 0 || x >= n || y >= m || lines[y][x] != c.val+1 {
						continue
					}

					if lines[y][x] == '9' {
						complete([2]int{i, j}, [2]int{y, x})
					} else {
						q.Push(cell{lines[y][x], x, y})
					}
				}
			}
		}
	}
}

func part1(lines []string) int {
	seen := map[[2]position]empty{}

	walk(lines, func(src, dst position) {
		seen[[2]position{src, dst}] = empty{}
	})

	return len(seen)
}

func part2(lines []string) int {
	count := 0

	walk(lines, func(src, dst position) {
		count++
	})

	return count
}
