package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/lluchkaa/aoc/internal/reader"
)

const filename = "input/day11.txt"

func main() {
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		panic(err)
	}

	line := reader.ReadLines(file)[0]

	parts := strings.Split(line, " ")

	stones := make(map[int]int)

	for _, p := range parts {
		stone, _ := strconv.Atoi(p)
		if _, ok := stones[stone]; !ok {
			stones[stone] = 0
		}
		stones[stone]++
	}

	for i := 0; i < 75; i++ {
		if i == 25 {
			fmt.Printf("Part 1: %d\n", count(stones))
		}
		stones = blink(stones)
	}
	fmt.Printf("Part 2: %d\n", count(stones))
}

func digits(n int) int {
	return int(math.Log10(float64(n)) + 1)
}

func transform(stone int) []int {
	if stone == 0 {
		return []int{1}
	}

	d := digits(stone)

	if d%2 == 0 {
		pow := int(math.Pow10(d / 2))
		return []int{stone / pow, stone % pow}
	}

	return []int{stone * 2024}
}

func blink(stones map[int]int) map[int]int {
	m := make(map[int]int)

	for k, v := range stones {
		transformed := transform(k)

		for _, t := range transformed {
			if _, ok := m[t]; !ok {
				m[t] = 0
			}
			m[t] += v
		}
	}

	return m
}

func count(stones map[int]int) int {
	sum := 0

	for _, v := range stones {
		sum += v
	}

	return sum
}
