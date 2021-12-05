package piscine

func IterativeFactorial(nb int) int {
	result1 := 1
	if nb > 0 {
		for i := 1; i < nb+1 ; i++ {
			result1 = result1 * i
		}
		return result1
	} else {
		return 0
	}

}
