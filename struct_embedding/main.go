package main

import "fmt"

type base struct {
	num int
}

type container struct {
	base
	str string
}

type describer interface {
	describe() string
}

func (b base) describe() string {
	return fmt.Sprintf("base with num = %v", b.num)
}

func main() {

	// embedded structs accessed with the name of their type
	co := container{
		base: base{
			num: 1,
		},
		str: "some name",
	}

	fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str)
	fmt.Printf("Also num: %v\n", co.base.num)
	fmt.Printf("describer:%v\n", co.describe())

	var desc describer = co

	fmt.Printf("describer:%v\n", desc.describe())

}
