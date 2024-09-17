package main

import "os"

// Panic is used for failing fast on errors we don't want to handle gracefully

func main() {

	// panic("Problem")

	_, err := os.Create("tempfile.go")
	if err != nil {
		panic(err)
	}

}
