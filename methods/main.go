package main

import (
	"fmt"
)

type rect struct {
	width, height int
}

/*
In Go, when you have a method with a pointer receiver, like func (r *rect) area() int, the compiler will automatically dereference the pointer for you when you access the fields of the rect struct.
Go doesn't need to use C++'s -> ; it implicitly dereferences with x.y
*/

func (r *rect) area() int {
	return r.width * r.height
}

func (r rect) peram() int {
	return 2*r.width + 2*r.height
}

// NB Implicit dereferencing applies to pointer recievers
// IT DOES NOT APPLY TO PARAMETERS
func (r *rect) scaleMe(s int) int {
	return s * r.area()
}

func main() {

	x := 5
	r := rect{width: 10, height: 20}

	r.scaleMe(x)
}
