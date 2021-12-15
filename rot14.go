package piscine

func Rot14(s string) string {
	r := []rune(s)

	for i := 0; i < len(s); i++ {
		if (r[i] >= 'a' && r[i] <= 'z') || (r[i] >= 'A' && r[i] <= 'Z') {
			if (r[i] <= 'L' && r[i] >= 'A') || (r[i] <= 'l' && r[i] >= 'a') {
				r[i] = r[i] + 14
			} else if r[i] <= 'Z' && r[i] >= 'M' {
				x := 90 - r[i]
				r[i] = 65 - x + 13
			} else if r[i] <= 'z' && r[i] >= 'm' {
				y := 122 - r[i]
				r[i] = 97 - y + 13
			}
		}
	}

	return string(r)
}
