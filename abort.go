package piscine

func Abort(a, b, c, d, e int) int {
	intSlice := []int{a, b, c, d, e}
	for i := 0; i < len(intSlice)-1; i++ {
		for j := 0; j < len(intSlice)-1; j++ {
			if intSlice[j] > intSlice[j+1] {
				intSlice[j], intSlice[j+1] = intSlice[j+1], intSlice[j]
			}
		}
	}
	return intSlice[2]
}
