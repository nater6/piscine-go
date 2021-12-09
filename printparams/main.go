package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args

	for i := 1; i < len(args); i++ {
		word := args[i]
		for _, letter := range word {
			z01.PrintRune(rune(letter))
		}
		z01.PrintRune('\n')
	}
}
