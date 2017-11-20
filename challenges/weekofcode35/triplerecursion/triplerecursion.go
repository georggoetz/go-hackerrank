// http://www.hackerrank.com/contests/w35/challenges/triple-recursion

package triplerecursion

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func Solve(r io.Reader, w io.Writer) {
	var n, m, k int
	fmt.Fscanf(r, "%d %d %d\n", &n, &m, &k)
	a := make([][]int, n)
	for i := range a {
		a[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		if i == 0 {
			a[i][i] = m
		} else {
			a[i][i] = a[i-1][i-1] + k
		}
		for j := i + 1; j < n; j++ {
			a[i][j] = a[i][j-1] - 1
			a[j][i] = a[j-1][i] - 1
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if j == n-1 {
				fmt.Fprintf(w, "%d\n", a[i][j])
			} else {
				fmt.Fprintf(w, "%d ", a[i][j])
			}
		}
	}
}

func main() {
	Solve(bufio.NewReader(os.Stdin), os.Stdout)
}
