package piscine

func Split(s, sep string) []string {
	separated := []string{}
	wordIndex := 0

	if len(s) >= len(sep) {
		for i := 0; i <= len(s)-len(sep); i++ {
			if sep == s[i:i+len(sep)] {
				if s[wordIndex:i] == "" {
					wordIndex = i + len(sep)
					continue
				}
				separated = append(separated, s[wordIndex:i])
				wordIndex = i + len(sep)
			} else if i == len(s)-len(sep) {
				separated = append(separated, s[wordIndex:])
			}
		}
	}
	return separated
}
