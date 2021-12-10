package piscine

func MakeRange(min, max int) []int {
	if max-min < 1 {
		return []int(nil)
	}
	slice := make([]int, max-min)

	for i := min; i < max; i++ {
		slice[i-min] = i
	}
	return slice
}
