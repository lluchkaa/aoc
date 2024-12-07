package main

import (
	"fmt"
	"os"

	"bufio"
	"io"
	"slices"
	"strconv"
	"strings"
)

const filename = "input/day5.txt"

func main() {
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		panic(err)
	}

	rules, lists := readData(file)

	fmt.Printf("Part 1: %d\n", part1(lists, rules))
	fmt.Printf("Part 2: %d\n", part2(lists, rules))
}

type empty = struct{}

func parsePages(list string) []int {
	pages := make([]int, 0)

	for _, part := range strings.Split(list, ",") {
		num, _ := strconv.Atoi(part)
		pages = append(pages, num)
	}

	return pages
}

func parseRule(rule string) (int, int) {
	parts := strings.Split(rule, "|")

	first, _ := strconv.Atoi(parts[0])
	second, _ := strconv.Atoi(parts[1])

	return first, second
}

func readData(r io.Reader) (map[int]map[int]empty, [][]int) {
	scanner := bufio.NewScanner(r)

	rules := make(map[int]map[int]empty)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}

		first, second := parseRule(line)

		if _, ok := rules[first]; !ok {
			rules[first] = make(map[int]empty)
		}
		rules[first][second] = empty{}
	}

	pageLists := make([][]int, 0)
	for scanner.Scan() {
		pageLists = append(pageLists, parsePages(scanner.Text()))
	}

	return rules, pageLists
}

func isSorted(list []int, rules map[int]map[int]empty) bool {
	visited := make([]int, 0, len(list))

	for _, x := range list {
		for _, v := range visited {
			if _, ok := rules[x][v]; ok {
				return false
			}
		}

		visited = append(visited, x)
	}

	return true
}

func part1(lists [][]int, rules map[int]map[int]empty) int {
	res := 0

	for _, list := range lists {
		if isSorted(list, rules) {
			res += list[len(list)/2]
		}
	}

	return res
}

func sortList(list []int, rules map[int]map[int]empty) []int {
	slices.SortFunc(list, func(x, y int) int {
		if _, ok := rules[x][y]; ok {
			return -1
		} else if _, ok := rules[y][x]; ok {
			return 1
		}

		return 0
	})

	return list
}

func part2(lists [][]int, rules map[int]map[int]empty) int {
	res := 0

	for _, list := range lists {
		if !isSorted(list, rules) {
			sorted := sortList(list, rules)
			res += sorted[len(sorted)/2]
		}
	}

	return res
}
