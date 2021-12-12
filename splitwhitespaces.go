package piscine

func SplitWhiteSpaces(s string) []string {
	finalSlice := []string{}
	wordIndex := 0

	for i, letter := range s {
		if letter == ' ' || letter == '\t' || letter == '\n' {
			if s[wordIndex:i] == "" {
				wordIndex = i + 1
				continue
			}
			finalSlice = append(finalSlice, s[wordIndex:i])
			wordIndex = i + 1
		} else if i == len(s)-1 {
			finalSlice = append(finalSlice, s[wordIndex:])
		}
	}
	return finalSlice
}
