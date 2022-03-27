package documentMe

const Pie = 3.1415912

func S1(s string) int {
	if s == "" {
		return 0
	}
	n := 0
	for range s {
		n++
	}
	return n
}

func F1(n int) int {
	return 2 * n
}
