package piscine

func Map(f func(int) bool, a []int) []bool {
	boolSlice := []bool{}

	for i := 0; i < len(a); i++ {
		x := f(a[i])
		boolSlice = append(boolSlice, x)
	}
	return boolSlice
}
