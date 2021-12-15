package piscine

func Rot14(s string) string {
	for _, letter := range s {
		letter = letter + 14
	}

	return s
}
