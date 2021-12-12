package main

import (
	"os"

	"github.com/01-edu/z01"
)

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func isEven(nbr int) bool {
	if nbr%2 == 0 {
		printStr("I have an even number of arguments")
		return true
	} else {
		printStr("I have an odd number of arguments	")
		return false
	}
}

func main() {
	args := os.Args
	lengthOfArg := len(args) - 1
	isEven(lengthOfArg)
}
