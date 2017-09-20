package pairs

import (
	"fmt"
	"strings"
)

const (
	sample1 = `5 2
    1 5 3 4 2`

	sample2 = `10 1
    363374326 364147530 61825163 1073065718 1281246024 1399469912 428047635 491595254 879792181 1069262793`
)

func ExampleSolve() {
	fmt.Println(Solve(read(strings.NewReader(sample1))))
	fmt.Println(Solve(read(strings.NewReader(sample2))))
	// Output:
	// 3
	// 0
}
