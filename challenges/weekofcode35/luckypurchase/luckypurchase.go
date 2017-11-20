// http://www.hackerrank.com/contests/w35/challenges/lucky-purchase

package luckypurchase

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
)

type laptop struct {
	name  string
	price int
}

type byPrice []laptop

func (p byPrice) Len() int           { return len(p) }
func (p byPrice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p byPrice) Less(i, j int) bool { return p[i].price < p[j].price }

func Solve(r io.Reader, w io.Writer) {
	var (
		n, four, seven int
		name, price    string
		discard        bool
	)
	fmt.Fscanf(r, "%d\n", &n)
	a := make([]laptop, 0)
	for i := 0; i < n; i++ {
		fmt.Fscanf(r, "%s %s\n", &name, &price)
		four, seven, discard = 0, 0, false
		for _, r := range price {
			switch r {
			case '4':
				four++
			case '7':
				seven++
			default:
				discard = true
			}
		}
		if discard {
			continue
		}
		if four == seven {
			p, _ := strconv.Atoi(price)
			a = append(a, laptop{name: name, price: p})
		}
	}
	if len(a) == 0 {
		fmt.Fprintf(w, "%d\n", -1)
	} else {
		sort.Sort(byPrice(a))
		fmt.Fprintf(w, "%s\n", a[0].name)
	}
}

func main() {
	Solve(bufio.NewReader(os.Stdin), os.Stdout)
}
