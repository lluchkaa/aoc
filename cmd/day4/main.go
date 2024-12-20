package main

import (
	"fmt"
	"os"

	"github.com/lluchkaa/aoc/internal/reader"
)

const filename = "input/day4.txt"

func main() {
	inputFile, err := os.Open(filename)
	defer inputFile.Close()
	if err != nil {
		panic(err)
	}

	lines := reader.ReadLines(inputFile)
	count := getCountOfOccurrences(lines, "XMAS")
	xCount := getCountOfXMases(lines)

	fmt.Printf("Count: %d; X Count: %d\n", count, xCount)
}

type direction = [2]int

var directions = [...]direction{
	{0, -1},
	{1, -1},
	{1, 0},
	{1, 1},
	{0, 1},
	{-1, 1},
	{-1, 0},
	{-1, -1},
}

func search(text []string, x, y int, d direction, word string) bool {
	for idx := 0; idx < len(word); idx++ {
		if x < 0 || y < 0 || y >= len(text) || x >= len(text[y]) {
			return false
		}

		if text[y][x] != word[idx] {
			return false
		}

		x += d[0]
		y += d[1]
	}

	return true
}

func getCountOfOccurrences(lines []string, word string) int {
	res := 0
	for x := range lines {
		for y := range lines[x] {
			for _, dir := range directions {
				if search(lines, x, y, dir, word) {
					res++
				}
			}
		}
	}

	return res
}

func isXMAS(text []string, x, y int) bool {
	if x <= 0 || y <= 0 || y >= len(text)-1 || x >= len(text[y])-1 {
		return false
	}

	if text[y][x] != 'A' {
		return false
	}

	return text[y-1][x-1] == 'M' && text[y-1][x+1] == 'M' && text[y+1][x+1] == 'S' && text[y+1][x-1] == 'S' ||
		text[y-1][x-1] == 'S' && text[y-1][x+1] == 'M' && text[y+1][x+1] == 'M' && text[y+1][x-1] == 'S' ||
		text[y-1][x-1] == 'S' && text[y-1][x+1] == 'S' && text[y+1][x+1] == 'M' && text[y+1][x-1] == 'M' ||
		text[y-1][x-1] == 'M' && text[y-1][x+1] == 'S' && text[y+1][x+1] == 'S' && text[y+1][x-1] == 'M'
}

func getCountOfXMases(lines []string) int {
	res := 0
	for y := range lines {
		for x := range lines[y] {
			if isXMAS(lines, x, y) {
				res++
			}
		}
	}

	return res
}
