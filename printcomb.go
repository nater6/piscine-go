package piscine

import "github.com/01-edu/z01"

func PrintComb() {
	var i rune = '0'
	var j rune = '0'
	var k rune = '0'

	for i <= 9 {
		if i < j && (j < k) {
			z01.PrintRune(i)
			z01.PrintRune(j)
			z01.PrintRune(k)
			z01.PrintRune(',')
			z01.PrintRune(' ')
		} else if i >= j {
			j++
		} else if j >= k {
			k++
		} else {
			i++
			j = 0
			k = 0
		}
	}
}
