package math

// MinInt returns the smallest of v...vn.
func MinInt(v int, vn ...int) int {
	m := v
	for _, x := range vn {
		if x < m {
			m = x
		}
	}
	return m
}

// MaxInt returns the largest of v...vn.
func MaxInt(v int, vn ...int) int {
	m := v
	for _, x := range vn {
		if x > m {
			m = x
		}
	}
	return m
}

// AbsInt returns the absolute value of v.
func AbsInt(v int) int {
	if v < 0 {
		return -v
	}
	return v
}
