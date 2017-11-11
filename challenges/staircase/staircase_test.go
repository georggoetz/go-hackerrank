package staircase

import (
	"fmt"
	"os"
	"strings"
)

func ExampleSolve() {
	fmt.Println()
	Solve(strings.NewReader("4"), os.Stdout)
	// Output:
	//    #
	//   ##
	//  ###
	// ####
}
