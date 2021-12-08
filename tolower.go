package piscine

func ToLower(s string) string {
	r := []rune(s)
	for i := 0; i < len(s); i++ {
		if r[i] >= 'A' && r[i] <= 'Z' {
			r[i] = r[i] + 32
		}
	}
	return string(r)
}
