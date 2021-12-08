package piscine

func IsLower(s string) bool {
	rune1 := []rune(s)
	for i := 0; i < len(s); i++ {
		if rune1[i] > 'z' || rune1[i] < 'a' {
			return false
		}
	}
	return true
}
