package piscine

func NRune(s string, n int) rune {
	if n <= len(s)-1 {
		rune1 := []rune(s)
		return rune1[n]
	} else {
		return 0
	}
}
