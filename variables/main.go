package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a = "variable"
	var b rune = 'A'
	var aType = reflect.TypeOf(a)

	fmt.Println(aType)
	fmt.Println(b)
}
