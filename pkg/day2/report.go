package day2

import (
	"strconv"
	"strings"
)

type report = []int

func parseLine(s string) report {
	parts := strings.Split(s, " ")
	report := make(report, len(parts))

	for i, x := range parts {
		report[i], _ = strconv.Atoi(x)
	}

	return report
}

func isAscSafe(x, y int) bool {
	return y > x && y-x <= 3
}

func isDescSafe(x, y int) bool {
	return x > y && x-y <= 3
}

func isSafeFunc(r report, comparator func(x, y int) bool) bool {
	if len(r) < 2 {
		return true
	}

	for i := 1; i < len(r); i++ {
		if !comparator(r[i-1], r[i]) {
			return false
		}
	}

	return true
}

func isSafe(r report) bool {
	if len(r) < 2 {
		return true
	}

	comparator := isAscSafe
	if r[0] > r[1] {
		comparator = isDescSafe
	}

	return isSafeFunc(r, comparator)
}

func isSafeWithDampenerFunc(r report, comparator func(x, y int) bool) bool {
	if len(r) < 2 {
		return true
	}

	for i := 1; i < len(r); i++ {
		if !comparator(r[i-1], r[i]) {
			if isSafeFunc(append([]int{r[i-1]}, r[i+1:]...), comparator) {
				return true
			}
			withoutPrevious := []int{}
			if i >= 2 {
				withoutPrevious = append(withoutPrevious, r[i-2:i-1]...)
			}
			withoutPrevious = append(withoutPrevious, r[i:]...)
			return isSafeFunc(withoutPrevious, comparator)
		}
	}

	return true
}

func isSafeWithDampenerBruteForceFunc(r report, comparator func(x, y int) bool) bool {
	arr := make([]int, len(r)-1)

	for i := 0; i < len(r); i++ {
		copy(arr[:i], r[:i])
		copy(arr[i:], r[i+1:])
		if isSafeFunc(arr, comparator) {
			return true
		}
	}

	return false
}

func isSafeWithDampener(r report) bool {
	return isSafe(r) || isSafeWithDampenerFunc(r, isAscSafe) || isSafeWithDampenerFunc(r, isDescSafe)
}
