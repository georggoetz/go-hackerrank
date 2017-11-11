package countingsort

func Solve2(a []int) (s []int) {
	s = make([]int, len(a))
	c := Solve1(a)
	si := 0
	for i, n := range c {
		for j := 0; j < n; j++ {
			s[si] = i
			si++
		}
	}
	return
}
