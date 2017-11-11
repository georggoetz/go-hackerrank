package countingsort

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
)

func Solve(r io.Reader, w io.Writer) {
	var (
		n, x int
		s    string
		b    bytes.Buffer
	)
	scanner := bufio.NewScanner(r)
	scanner.Scan()
	fmt.Sscanf(scanner.Text(), "%d", &n)
	a := make([]string, n)
	c := make([][]int, 100)
	for i := range c {
		c[i] = make([]int, 0)
	}
	for i := 0; i < n; i++ {
		scanner.Scan()
		fmt.Sscanf(scanner.Text(), "%d %s", &x, &s)
		if i < n/2 {
			a[i] = "-"
		} else {
			a[i] = s
		}
		c[x] = append(c[x], i)
	}
	for i := range c {
		for j := range c[i] {
			b.WriteString(a[c[i][j]])
			b.WriteString(" ")
		}
	}
	fmt.Fprint(w, b.String())
}
