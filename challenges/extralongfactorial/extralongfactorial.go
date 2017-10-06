// http://www.hackerrank.com/challenges/extra-long-factorials

package extralongfactorial

import (
	"fmt"
	"io"
	"math/big"
)

var (
	zero = big.NewInt(0)
	one  = big.NewInt(1)
)

// Factorial returns n!
func Factorial(n *big.Int) (f *big.Int) {
	f = new(big.Int)
	switch n.Cmp(zero) {
	case -1, 0:
		f.SetInt64(1)
	default:
		f.Set(n)
		f.Mul(f, Factorial(n.Sub(n, one)))
	}
	return
}

func read(r io.Reader) (n int) {
	fmt.Fscanf(r, "%d\n", &n)
	return
}
