package piscine

func Capitalize(s string) string {
	r := []rune(s)
	l := 0

	for i := 0; i < len(s); i++ {
		if l == 0 && (r[i] >= 'a' && r[i] <= 'z') {
			r[i] = r[i] - 32
			l++
		} else if l == 0 && ((r[i] >= 'A' && r[i] <= 'Z') || (r[i] >= '0' && r[i] <= '9')) {
			l++
		} else if (r[i] >= 'A' && r[i] <= 'Z') {
			r[i] = r[i] + 32
			l++
		} else if (r[i] >= '0' && r[i] <= '9') || (r[i] >= 'a' && r[i] <= 'z')  {
			l++
		} else {
			l=0
		} 
	}
	return string(r)
}
