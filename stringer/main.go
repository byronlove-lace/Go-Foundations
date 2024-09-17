package main

// stringer is used for custom string implementation of datatypes

import "fmt"

type Person struct {
	Name    string
	Age     int
	Hobbies []string
}

func (p Person) String() string {
	return "Hello"
}

func main() {

	p1 := Person{"Sam", 20, []string{"cricket", "football"}}
	fmt.Println(p1)
}
