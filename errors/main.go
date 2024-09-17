package main

import (
	"errors"
	"fmt"
)

func funshion(arg int) (int, error) {

	if arg == 42 {
		return -1, errors.New("42? What the fuck is wrong with you!?")
	}

	return arg + 3, nil
}

var ErrOutOfTea = fmt.Errorf("No more tea")
var ErrNoPower = fmt.Errorf("No more Power")

func makeTea(arg int) error {
	if arg == 2 {
		return ErrOutOfTea
	} else if arg == 4 {
		return fmt.Errorf("making tea: %w", ErrNoPower)
	}

	return nil
}

func main() {

	result, err := funshion(4)

	// common pattern if err != nil
	if result == -1 {
		// Sprintf returns string Printf outputs to screen
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	nuErr := makeTea(4)
	fmt.Println(nuErr)

}
