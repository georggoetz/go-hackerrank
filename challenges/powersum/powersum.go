// http://www.hackerrank.com/challenges/the-power-sum

package powersum

import (
	"fmt"
	"io"
	"math"
)

// Solve finds the number of ways an integer n can be expressed as sum of k-th
// power of unique numbers.
func Solve(n, x int) int {
	return powerSum(n, x, 1)
}

func powerSum(n, x, i int) int {
	v := n - int(math.Pow(float64(i), float64(x)))
	if v < 0 {
		return 0
	}
	if v == 0 {
		return 1
	}
	return powerSum(v, x, i+1) + powerSum(n, x, i+1)
}

func read(r io.Reader) (x, n int) {
	fmt.Fscanf(r, "%d\n%d\n", &x, &n)
	return
}
