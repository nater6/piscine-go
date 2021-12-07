package piscine

func AlphaCount(s string) int {
	rune1 := []rune(s)

	i := 0
	for i < len(s) {
		if (rune1[i] >= 'a' && rune1[i] <= 'z') || (rune1[i] >= 'A' && rune1[i] <= 'Z') {
			i++
		}
	}
	return i
}
