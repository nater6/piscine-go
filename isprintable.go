package piscine

func IsPrintable(s string) bool {
	r := []rune(s)

	for i := 0; i < len(s); i++ {
		if r[i] > 32 || r[i] < 127 {
		} else {
			return false
		}
	}
	return true
}
