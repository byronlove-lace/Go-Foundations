package main

import "fmt"

func main() {

	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	for j := range 2 {
		fmt.Println("RANGE:", j)
	}
}
