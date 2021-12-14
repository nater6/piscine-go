package main

import (
	"github.com/01-edu/z01"
)

type pointers struct {
	x string
	y string
}

func main() {
	points := pointers{}
	points.x = "42"
	points.y = "21"

	s := "x = " + points.x + ", y = " + points.y

	for _, letters := range s {
		z01.PrintRune(letters)
	}
	z01.PrintRune('\n')
}
