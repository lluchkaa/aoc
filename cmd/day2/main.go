package main

import (
	"fmt"
	"os"

	"github.com/lluchkaa/aoc/pkg/day2"
)

const filename = "input/day2.txt"

func main() {
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Safe reports count: %d\n", day2.GetCountOfSafeReports(file))
	file.Seek(0, 0)

	fmt.Printf("Safe reports with Dampener count: %d\n", day2.GetCountOfSafeReportsWithDampener(file))
}
