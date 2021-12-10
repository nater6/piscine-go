package piscine

func AppendRange(min, max int) []int {
	slice := []int{}

	for i := min; i <= max; i++ {
		slice = append(slice, i)
	}
	return slice
}
