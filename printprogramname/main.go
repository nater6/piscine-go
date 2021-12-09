package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguements := os.Args
	name := arguements[0]

	for i, l := range name {
		if i > 1 {
			z01.PrintRune(rune(l))
		}
	}
	z01.PrintRune('\n')
}
