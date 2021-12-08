package piscine

func IsAlpha(s string) bool {
	rune1 := []rune(s)

	for i := 0; i < len(s); i++ {
		if (rune1[i] < 'a' || (rune1[i] > 'z' && rune1[i] < 'A') || rune1[i] > 'Z') && rune1[i] != ' ' {
			return false
		}
	}

	return true
}
