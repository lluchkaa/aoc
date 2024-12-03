package main

import (
	"fmt"
	"os"

	"github.com/lluchkaa/aoc/pkg/day3"
)

const filename = "input/day3.txt"

func main() {
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		panic(err)
	}

	sum, sumWithEnabled := day3.GetMultipliesSum(file)

	fmt.Printf("Sum: %d, Sum with enabled: %d\n", sum, sumWithEnabled)
}
