package main

import (
	"fmt"
	"reflect"
	"time"
)

func main() {

	i := 2
	fmt.Println("check", i, 4)
	fmt.Print("check\n")

	t := time.Now()
	hour := t.Hour
	var hType = reflect.TypeOf(hour)

	fmt.Println(hType)

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")

}
