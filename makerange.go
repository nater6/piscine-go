package piscine

func MakeRange(min, max int) []int {
	if max <= min {
		return []int(nil)
	}
	slice := make([]int, max-min)

	for i := min; i < max; i++ {
		slice[i] = i
	}
	return slice
}
