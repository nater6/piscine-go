package piscine

func Rot14(s string) string {
	r := []rune(s)

	for i := 0; i < len(s); i++ {
		if (r[i] >= 'a' && r[i] <= 'z') || (r[i] >= 'A' && r[i] <= 'Z') {
			if (r[i] <= 76 && r[i] >= 65) || (r[i] <= 108 && r[i] >= 97) {
				r[i] = r[i] + 14
			} else if r[i] <= 90 && r[i] >= 77 {
				x := 90 - r[i]
				r[i] = 65 - x + 14
			} else if r[i] <= 122 && r[i] >= 109 {
				y := 122 - r[i]
				r[i] = 97 - y + 14
			}
		}
	}

	return string(r)
}
