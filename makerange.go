package piscine

func MakeRange(min, max int) []int {
	slice := make([]int, max-min)

	if max <= min {
		return []int(nil)
	}
	for i := min; i < max; i++ {
		slice[i] = i
	}
	return slice
}
