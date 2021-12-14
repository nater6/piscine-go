package piscine

func Map(f func(int) bool, a []int) []bool {
	boolSlice := make([]bool, len(a))

	for i := 0; i < len(a); i++ {
		boolSlice = append(boolSlice, f(a[i]))
	}
	return boolSlice
}
