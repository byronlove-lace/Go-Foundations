package main

import "fmt"

func main() {
	var kph = 50

	if mph := kph * 1000; mph > 40000 {
		fmt.Println("WE DID IT")
	} else {
		fmt.Println("WE FAILED")
	}
}
