package main

import (
	"fmt"
)

func main() {

	nums := []int{1, 2, 3, 4}
	sum := 0

	// first value is the index
	for _, num := range nums {
		sum += num
	}

	fmt.Println(sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println(i)
		}
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	for k := range kvs {
		fmt.Printf("%s\n", k)
	}

}
