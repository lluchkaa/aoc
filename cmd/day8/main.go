package main

import (
	"fmt"
	"github.com/lluchkaa/aoc/internal/reader"
	"io"
	"os"
)

const filename = "input/day8.txt"

func main() {
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		panic(err)
	}

	cols, rows := reader.Dimensions(file)
	file.Seek(0, 0)

	part1Count := Solve(file, cols, rows, part1)
	fmt.Printf("Part 1: %d\n", part1Count)
	file.Seek(0, 0)

	part2Count := Solve(file, cols, rows, part2)
	fmt.Printf("Part 2: %d\n", part2Count)

}

type empty = struct{}
type position = [2]int
type frequency = byte

func part1(antinodes map[position]empty, pos1, pos2 position, w, h int) {
	dx, dy := pos1[0]-pos2[0], pos1[1]-pos2[1]

	if x, y := pos2[0]-dx, pos2[1]-dy; x >= 0 && y >= 0 && x < w && y < h {
		antinodes[position{x, y}] = empty{}
	}
	if x, y := pos1[0]+dx, pos1[1]+dy; x >= 0 && y >= 0 && x < w && y < h {
		antinodes[position{x, y}] = empty{}
	}
}

func part2(antinodes map[position]empty, pos1, pos2 position, w, h int) {
	dx, dy := pos1[0]-pos2[0], pos1[1]-pos2[1]

	for x, y := pos1[0], pos1[1]; x >= 0 && y >= 0 && x < w && y < h; x, y = x-dx, y-dy {
		antinodes[position{x, y}] = empty{}
	}
	for x, y := pos1[0], pos1[1]; x >= 0 && y >= 0 && x < w && y < h; x, y = x+dx, y+dy {
		antinodes[position{x, y}] = empty{}
	}
}

func Solve(r io.Reader, w, h int, updateAntinodes func(map[position]empty, position, position, int, int)) int {
	m := make(map[frequency][]position)
	antinodes := make(map[position]empty)

	for y, line := range reader.Lines2(r) {
		for x, ch := range []byte(line) {
			if ch == '.' {
				continue
			}

			if _, ok := m[ch]; !ok {
				m[ch] = []position{}
			}

			for _, pos := range m[ch] {
				updateAntinodes(antinodes, pos, position{x, y}, w, h)
			}

			m[ch] = append(m[ch], position{x, y})
		}
	}

	return len(antinodes)
}
