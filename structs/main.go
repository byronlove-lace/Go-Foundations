package main

import (
	"fmt"
	"reflect"
)

// type is basically typedef
// using lowercase for structs like a pleb
type person struct {
	name string
	age  int
}

// Sophisticated constructor
func newPerson(name string) *person {

	p := person{name: name}
	p.age = 43
	return &p
}

func main() {
	var P1 *person = newPerson("KT")
	P2 := *newPerson("KT")

	// omitted fieds are zero valued
	fmt.Println(person{"bob", 20})
	fmt.Println(reflect.TypeOf(P1))
	fmt.Println(reflect.TypeOf(P2))
	fmt.Println(P2.name)
	fmt.Println(P2.age)

	// Initialization pattern
	// Only works for anonymous structs
	dog := struct {
		name   string
		isGood bool
	}{
		"Rex",
		true,
	}

	fmt.Println(dog)
}
