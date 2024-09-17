package main

import (
	"fmt"
	"reflect"
)

func addN(x *int) {
	*x++
}

func main() {
	const s string = "constant"

	var n int = 5

	const d = 3e20

	var dType = reflect.TypeOf(d)

	addN(&n)

	fmt.Println(n)
	fmt.Println(dType)
	fmt.Println(d)
}
