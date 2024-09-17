package main

import "fmt"

func main() {
	i := "hello"

	s := i.(int)

	fmt.Println(s)
}
