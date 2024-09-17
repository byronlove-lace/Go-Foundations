package main

import "fmt"

func main() {

	var a [5]int

	fmt.Println("emp ", a) // prints the whole array
	fmt.Println("emp ", len(a))

	b := [7]int{1, 2, 4: 3, 4, 5} // x: will set the index; all after
	// var c [5]int = [5]int{1, 2, 3, 4, 5}
	// d := [...]rune{'a', 4: 'b', 'c'}

	fmt.Println(b)

}
