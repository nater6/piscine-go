package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	sortedTester := []int{}
	for i := 0; i < len(a)-1; i++ {
		if f(a[i], a[i+1]) > 0 {
			sortedTester = append(sortedTester, 0)
		} else if f(a[i], a[i+1]) < 0 {
			sortedTester = append(sortedTester, 1)
		}
	}
	for j := 0; j < len(sortedTester)-1; j++ {
		if sortedTester[j] != sortedTester[j+1] {
			return false
		}
	}
	return true
}
