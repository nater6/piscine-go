package piscine

func Capitalize(s string) string {
	r := []rune(s)

	for i := 0; i < len(s); i++ {
		if r[i] == ' ' {
			if r[i+1] >= 'a' && r[i] <= 'z' {
				r[i+1] -= 32
			}
			for j := i; r[j] == ' '; j++ {
				if r[j] >= 'A' && r[j] <= 'Z' {
					r[j] += 32
				}
			}
		}
	}
	return string(r)
}
