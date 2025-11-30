package recursion

// FibN is not recursion function, but actually better
func FibN(n int) int {
	if n == 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	a, b := 0, 1
	for i := 0; i < n-1; i++ {
		temp := b
		b = a + b
		a = temp
	}
	return b
}
