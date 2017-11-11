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

func MinIntSlice(v []int) (m int) {
	if len(v) > 0 {
		m = v[0]
	}
	for i := 1; i < len(v); i++ {
		if v[i] < m {
			m = v[i]
		}
	}
	return
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

func MaxIntSlice(v []int) (m int) {
	if len(v) > 0 {
		m = v[0]
	}
	for i := 1; i < len(v); i++ {
		if v[i] > m {
			m = v[i]
		}
	}
	return
}

// AbsInt returns the absolute value of v.
func AbsInt(v int) int {
	if v < 0 {
		return -v
	}
	return v
}

// PowInt computes a^b
func PowInt(a, b int) (r int) {
	r = 1
	for b > 0 {
		if b&1 != 0 {
			r *= a
		}
		b >>= 1
		a *= a
	}
	return
}
