package piscine

func AppendRange(min, max int) []int {
	slice := []int{}

	if max <= min {
		return []int(nil)
	}

	for i := min; i < max; i++ {
		slice = append(slice, i)
	}
	return slice
}
