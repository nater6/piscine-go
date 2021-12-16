package main

import (
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args

	if len(args) != 2 {
		return
	}

	x, _ := strconv.Atoi(args[1])
	if x == 1 {
		z01.PrintRune('t')
		z01.PrintRune('r')
		z01.PrintRune('u')
		z01.PrintRune('e')
		return
	}

	i := 2
	for i <= x {
		if i == x {
			z01.PrintRune('t')
			z01.PrintRune('r')
			z01.PrintRune('u')
			z01.PrintRune('e')
			return
		} else {
			i *= 2
		}
	}
	z01.PrintRune('f')
	z01.PrintRune('a')
	z01.PrintRune('l')
	z01.PrintRune('s')
	z01.PrintRune('e')
}
