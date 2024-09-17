package main

import (
	"fmt"
	"slices"
)

func main() {

	// Declaration of an array
	// var myArr [5]int
	// Declaration of a slice
	var mySlice []string
	// init on declaration
	// mySlice := []int{1, 2, 4: 3, 4, 5} // x: will set the index; all after

	mySlice = make([]string, 3)

	fmt.Println("emp:", mySlice, "len:", len(mySlice), "cap:", cap(mySlice))

	mySlice[0] = "First"
	mySlice[2] = "Last"

	// append always adds to the end
	// if end is occuppied: it expands the slice and adds to the end
	mySlice = append(mySlice, "NU LAST, JA!")

	fmt.Println("ful:", mySlice)

	kimboSlice := make([]string, len(mySlice))
	// Go maintaints C's strcpy(destination, source) syntax. Because nothing is perfect.
	// This is nonetheless, an awesome built-in
	copy(kimboSlice, mySlice)

	sonOfASlice := kimboSlice[0:3]
	fmt.Println("ful:", sonOfASlice)

	grandSonOfASlice := sonOfASlice[1:]
	fmt.Println("ful:", grandSonOfASlice)

	// THIS is what slices is: general slice functionality is builtin
	if slices.Equal(sonOfASlice, grandSonOfASlice) {
		fmt.Println("They're the same person!")
	}

	var mechaSlice = make([][]int, 3)

	for i := 0; i < 3; i++ {
		innerLen := i + 1
		mechaSlice[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			mechaSlice[i][j] = i + j
		}
	}

	fmt.Println("BIG:", mechaSlice)

	delSlice := []int{1, 2, 3, 4}
	var delInd = 2

	// Same approach used when we delete an element from an array
	// ... unpacks the elements of a slice
	delSlice = append(delSlice[:delInd], delSlice[delInd+1:]...)

}
