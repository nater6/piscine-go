package piscine

func Compact(ptr *[]string) int {
	a := *ptr
	nonZeroCounter := 0
	var newSlice []string

	for i := 0; i < len(a); i++ {
		if a[i] != "" {
			nonZeroCounter++
			newSlice = append(newSlice, a[i])
		}
	}

	*ptr = newSlice

	return nonZeroCounter
}
