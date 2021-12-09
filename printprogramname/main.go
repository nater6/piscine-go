package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	a := (os.Args[0])

	arune := []rune(a)

	for i := 0; i < len(arune); i++ {
		if arune[i] == '.' && arune[i+1] == '/' {
			fiojg := arune[i+2:]

			for _, j := range fiojg {
				z01.PrintRune(j)
			}
		}
	}
}
