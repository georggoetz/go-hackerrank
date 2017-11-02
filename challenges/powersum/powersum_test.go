package powersum

import "fmt"

func ExampleSolve() {
	fmt.Println(Solve(10, 2))
	fmt.Println(Solve(100, 2))
	fmt.Println(Solve(100, 3))
	// Output:
	// 1
	// 3
	// 1
}
