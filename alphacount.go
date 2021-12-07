package piscine

func AlphaCount(s string) int {
	rune1 := []rune(s)

	x := 0
	for i := 0; i < len(s); i++ {
		if (rune1[i] >= 'a' && rune1[i] <= 'z') || (rune1[i] >= 'A' && rune1[i] <= 'Z') {
			x++
		}
	}
	return x
}
