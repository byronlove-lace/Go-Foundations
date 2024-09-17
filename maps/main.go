package main

import (
	"fmt"
	"maps"
	// "maps"
)

func main() {

	m := make(map[string]int)

	m["e1"] = 1
	m["e2"] = 12

	fmt.Println(m)

	var e1 int = m["e1"]

	fmt.Println(e1)

	delete(m, "e1")
	fmt.Println(m)

	// second value is a bool on whether or not there is a hit
	// blank identifier _ is used when we don't want to use the first return value
	_, value := m["e3"]

	fmt.Println(value)

	// removes ALL k/v values
	clear(m)

	nm := map[string]int{"first": 1, "second": 2}
	nm2 := map[string]int{"first": 1, "second": 2}

	if maps.Equal(nm, nm2) {
		fmt.Println(nm)
	}

}
