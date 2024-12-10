package reader

import (
	"bufio"
	"io"
	"iter"
)

func Lines(r io.Reader) iter.Seq[string] {
	scanner := bufio.NewScanner(r)

	return func(yield func(string) bool) {
		for scanner.Scan() {
			if !yield(scanner.Text()) {
				break
			}
		}
	}
}

func Lines2(r io.Reader) iter.Seq2[int, string] {
	scanner := bufio.NewScanner(r)
	i := 0

	return func(yield func(int, string) bool) {
		for scanner.Scan() {
			if !yield(i, scanner.Text()) {
				break
			}
			i++
		}
	}
}
