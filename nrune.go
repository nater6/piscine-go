package piscine

func NRune(s string, n int) rune {
	if 0 <= n && n <= len(s)-1 {
		rune1 := []rune(s)
		return rune1[n-1]
	} else {
		return 0
	}
}
