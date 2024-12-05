package main

import (
	"fmt"
	"os"

	"github.com/lluchkaa/aoc/pkg/day5"
)

const filename = "input/day5.txt"

func main() {
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		panic(err)
	}

	rules, lists := day5.ReadData(file)

	fmt.Printf("Part 1: %d\n", day5.Part1(lists, rules))
	fmt.Printf("Part 2: %d\n", day5.Part2(lists, rules))
}
