package piscine

func IsAlpha(s string) bool {
	rune1 := []rune(s)

	for i := 0; i < len(s); i++ {
		if rune1[i] >= 'a' && rune1[i] <= 'z' {
		} else if rune1[i] >= 'A' && rune1[i] <= 'Z' {
		} else if rune1[i] == ' ' {
		} else if rune1[i] >= '0' && rune1[i] <= '9' {
		} else {
			return false
		}
	}

	return true
}
