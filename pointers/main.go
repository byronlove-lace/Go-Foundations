package main

import "fmt"

var num int = 5

func iter(iterme int) {
	iterme++
}

func betterIter(iterme *int) {
	*iterme++
}

func main() {
	iter(num)
	fmt.Println("dumb: ", num)

	betterIter(&num)
	fmt.Println("smart: ", num)

}
