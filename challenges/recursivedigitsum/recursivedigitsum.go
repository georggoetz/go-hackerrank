// http://www.hackerrank.com/challenges/recursive-digit-sum

package recursivedigitsum

import (
	"fmt"
	"io"
	"math/big"
)

//
func Solve(n string, k int64) string {
	ans := new(big.Int)
	ans.SetString(n, 10)
	ans.Mul(ans, big.NewInt(k))
	ans.Mod(ans, big.NewInt(9))
	if ans.Cmp(big.NewInt(0)) == 0 {
		return "9"
	}
	return ans.String()
}

func read(r io.Reader) (n string, k int64) {
	fmt.Fscanf(r, "%s %d\n", &n, &k)
	return
}
