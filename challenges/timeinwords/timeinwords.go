package timeinwords

import (
	"fmt"
	"io"
)

var words = map[int]string{
	1:  "one",
	2:  "two",
	3:  "three",
	4:  "four",
	5:  "five",
	6:  "six",
	7:  "seven",
	8:  "eight",
	9:  "nine",
	10: "ten",
	11: "eleven",
	12: "twelve",
	13: "thirteen",
	14: "fourteen",
	15: "quarter",
	16: "sixteen",
	17: "seventeen",
	18: "eighteen",
	19: "nineteen",
	20: "twenty",
	21: "twenty one",
	22: "twenty two",
	23: "twenty three",
	24: "twenty four",
	25: "twenty five",
	26: "twenty six",
	27: "twenty seven",
	28: "twenty eight",
	29: "twenty nine",
	30: "half"}

// Converts returns the time given by the numerals h and m into words.
func Convert(h, m int) string {
	if m == 0 {
		return words[h] + " o' clock"
	}
	s := "past"
	if m > 30 {
		m = 60 - m
		h++
		h %= 12
		s = "to"
	}
	if m == 15 || m == 30 {
		s = words[m] + " " + s
	} else if m == 1 {
		s = words[m] + " minute " + s
	} else {
		s = words[m] + " minutes " + s
	}
	return s + " " + words[h]
}

func read(r io.Reader) (h, m int) {
	fmt.Fscanf(r, "%d\n", &h)
	fmt.Fscanf(r, "%d\n", &m)
	return
}
