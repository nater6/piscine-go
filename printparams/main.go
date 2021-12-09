package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args

	for _, word := range args {
		for _, letter := range word {
			z01.PrintRune(rune(letter))
		}
	}
	z01.PrintRune('\n')
}
