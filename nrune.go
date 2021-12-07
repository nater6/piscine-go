package piscine

func NRune(s string, n int) rune {
	if n > 0 && n <= len(s) {
		rune1 := []rune(s)
		return rune1[n-1]
	} else {
		return 0
	}
}
