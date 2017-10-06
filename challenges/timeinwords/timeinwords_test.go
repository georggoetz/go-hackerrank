package timeinwords

import "fmt"

func ExampleConvert() {
	fmt.Println(Convert(5, 0))
	fmt.Println(Convert(5, 1))
	fmt.Println(Convert(5, 10))
	fmt.Println(Convert(5, 15))
	fmt.Println(Convert(5, 30))
	fmt.Println(Convert(5, 40))
	fmt.Println(Convert(5, 45))
	fmt.Println(Convert(5, 47))
	fmt.Println(Convert(5, 28))
	fmt.Println(Convert(12, 59))
	// Output:
	// five o' clock
	// one minute past five
	// ten minutes past five
	// quarter past five
	// half past five
	// twenty minutes to six
	// quarter to six
	// thirteen minutes to six
	// twenty eight minutes past five
	// one minute to one
}
