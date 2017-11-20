// http://www.hackerrank.com/contests/w35/challenges/3d-surface-area

package surfacearea3d

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type coord struct {
	cubes, area int
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Solve(r io.Reader, w io.Writer) {
	var rows, cols, d int
	fmt.Fscanf(r, "%d %d\n", &rows, &cols)
	a := make([][]coord, rows)
	for i := range a {
		a[i] = make([]coord, cols)
		for j := 0; j < cols-1; j++ {
			fmt.Fscanf(r, "%d", &a[i][j].cubes)
		}
		fmt.Fscanf(r, "%d\n", &a[i][cols-1].cubes)
	}
	sum := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			a[i][j].area = 4*a[i][j].cubes + 2
			d = 0
			if j < cols-1 {
				d += minInt(a[i][j].cubes, a[i][j+1].cubes)
			}
			if j > 0 {
				d += minInt(a[i][j].cubes, a[i][j-1].cubes)
			}
			if i < rows-1 {
				d += minInt(a[i][j].cubes, a[i+1][j].cubes)
			}
			if i > 0 {
				d += minInt(a[i][j].cubes, a[i-1][j].cubes)
			}
			a[i][j].area -= d
			sum += a[i][j].area
		}
	}
	fmt.Fprintf(w, "%d\n", sum)
}

func main() {
	Solve(bufio.NewReader(os.Stdin), os.Stdout)
}
