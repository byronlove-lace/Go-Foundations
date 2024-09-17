package main

import (
	"fmt"
)

func iterSeq() func() int {
	i := 0
	// anonymous funcs
	return func() int {
		i++
		return i
	}
}

// return a function that returns an int
func cubd(base int) func() int {
	exponent := base * base

	return func() int {
		return exponent * exponent
	}
}

func main() {
	x := iterSeq()
	y := x()
	fmt.Println(y)

	exp := cubd(4)
	exp3 := exp()
	fmt.Println(exp3)
}
