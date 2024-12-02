package slice

func Frequency[T comparable](s []T) map[T]int {
	m := make(map[T]int)

	for _, x := range s {
		m[x] = m[x] + 1
	}

	return m
}
