package piscine

func IsUpper(s string) bool {
	rune1 := []rune(s)
	for i:=0; i<len(s); i++ {
		if rune1[i] > 'Z' || rune1[i] < 'A' {
			return false
		}
	}
	return true
}

