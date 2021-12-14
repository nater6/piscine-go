package piscine

func Countif(f func(string) bool, tab []string) int {
	trueCounter := []int{}

	for i := 0; i < len(tab); i++ {
		if f(tab[i]) {
			trueCounter = append(trueCounter, 0)
		}
	}

	return len(trueCounter)
}
