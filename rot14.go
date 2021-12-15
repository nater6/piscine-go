package piscine

import "github.com/01-edu/z01"

func Rot14(s string) string {
	for _, letter := range s {
		letter = letter + 14
	}

	return s
}
