package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args

	for i, j := 0, len(args)-1; i < j; i, j = i+1, j-1 {
		x := args[i]
		args[i] = args[j]
		args[j] = x
	}

	for a := 0; a < len(args)-1; a++ {
		word := args[a]
		for _, letter := range word {
			z01.PrintRune(rune(letter))
		}
		z01.PrintRune('\n')

	}
}
