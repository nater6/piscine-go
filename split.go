package piscine

func Split(s, sep string) []string {
	separated := []string{}
	wordIndex := 0

	if len(s) >= len(sep) {
		for i := 0; i < len(s); i++ {
			if sep == s[i:i+len(sep)] {
				if s[wordIndex:i] == "" {
					wordIndex = i + 1
					continue
				}
				separated = append(separated, s[wordIndex:i])
				wordIndex = i + 1
			} else if i == len(s)-1 {
				separated = append(separated, s)
			}
		}
	}
	return separated
}
