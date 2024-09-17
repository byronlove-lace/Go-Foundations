package main

import (
	"fmt"
)

func plus(a int, b int) int {
	return a + b
}

// type at end of list
func plusPlus(a, b, c int) int {
	return a + b + c
}

// NB go infers single return types but doesn't infer multiple returns
// need to explicitly declare
func vals() (int, int) {
	return 1, 2
}

// variadic functions
// NEED to declare return types
// equivalent to []int
func product(nums ...int) int {
	total := 0

	for i, num := range nums {
		if i == 0 {
			total += num
		} else {
			total *= num
		}
	}

	return total
}

func main() {

	x := plusPlus(1, 2, 3)
	fmt.Println(x)

	y := []int{1, 2, 3, 4, 5}
	prod := product(y...)
	fmt.Println(prod)
}
