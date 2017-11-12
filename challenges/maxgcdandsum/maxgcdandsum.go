// http://www.hackerrank.com/contests/w34/challenges/maximum-gcd-and-sum

package maxgcdandsum

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func Solve(r io.Reader, w io.Writer) {
	var n int
	fmt.Fscanf(r, "%d\n", &n)
	a := make([]int, n)
	b := make([]int, n)
	max := 0
	for i := 0; i < n; i++ {
		if i == n-1 {
			fmt.Fscanf(r, "%d\n", &a[i])
		} else {
			fmt.Fscanf(r, "%d", &a[i])
		}
		if a[i] > max {
			max = a[i]
		}
	}
	for i := 0; i < n; i++ {
		if i == n-1 {
			fmt.Fscanf(r, "%d\n", &b[i])
		} else {
			fmt.Fscanf(r, "%d", &b[i])
		}
		if b[i] > max {
			max = b[i]
		}
	}
	inA := make([]bool, max+1)
	inB := make([]bool, max+1)
	for i := 0; i < n; i++ {
		inA[a[i]] = true
		inB[b[i]] = true
	}
	mulA := make([]int, max+1)
	mulB := make([]int, max+1)
	for i := 1; i <= max; i++ {
		for j := i; j <= max; j += i {
			if inA[j] {
				mulA[i] = j
			}
			if inB[j] {
				mulB[i] = j
			}
		}
	}
	j := 0
	for i := 1; i <= max; i++ {
		if mulA[i] > 0 && mulB[i] > 0 {
			j = i
		}
	}
	fmt.Fprintf(w, "%d\n", mulA[j]+mulB[j])
}

func main() {
	Solve(bufio.NewReader(os.Stdin), os.Stdout)
}
