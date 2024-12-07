package main

import (
	"fmt"
	"os"

	"github.com/lluchkaa/aoc/pkg/day6"
)

const filename = "input/day6.txt"

func main() {
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		panic(err)
	}

	rules, lists := day6.ReadData(file)

	fmt.Printf("Part 1: %d\n", day6.Part1(lists, rules))
	fmt.Printf("Part 2: %d\n", day6.Part2(lists, rules))
}
