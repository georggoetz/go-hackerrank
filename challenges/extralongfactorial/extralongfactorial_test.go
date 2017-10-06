package extralongfactorial

import (
	"fmt"
	"math/big"
)

func ExampleFactorial() {
	fmt.Println(Factorial(big.NewInt(25)).String())
	// Output:
	// 15511210043330985984000000
}
