package main

import (
	"os"
	"strings"
)

func ExamplePairs() {
	testData := `5 2
    1 5 3 4 2`
	Pairs(strings.NewReader(testData), os.Stdout)

	testData = `10 1
    363374326 364147530 61825163 1073065718 1281246024 1399469912 428047635 491595254 879792181 1069262793`
	Pairs(strings.NewReader(testData), os.Stdout)

	// Output:
	// 3
	// 0
}
