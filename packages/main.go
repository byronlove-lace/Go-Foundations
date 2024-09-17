package main

// to import local packages w/o go mod init we use projectDir/packageDir
// rem: 1 package per dir or ur durrrrrrrrrr
import (
	"fmt"
	"packages/pkgs"
)

func main() {
	num := calculator.Add(1, 2)
	fmt.Println(num)
}
