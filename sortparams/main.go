package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	rSlice := []rune{}

	for i := 0; i < len(args); i++ {
		for _, r := range args[i] {
			rSlice = append(rSlice, r)
		}
	}

	for i := 0; i < len(rSlice); i++ {
		for j := i; j < len(rSlice); j++ {
			if rSlice[i] > rSlice[j] {
				rSlice[i], rSlice[j] = rSlice[j], rSlice[i]
			}
		}
	}

	for _, p := range rSlice {
		z01.PrintRune(p)
		z01.PrintRune('\n')
	}
}
