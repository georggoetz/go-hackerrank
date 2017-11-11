package countingsort

const max = 100

func Solve1(a []int) (c []int) {
	c = make([]int, max)
	for _, n := range a {
		c[n]++
	}
	return
}
