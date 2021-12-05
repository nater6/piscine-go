package piscine

func IterativeFactorial(nb int) int{
	result1 := 1
	if nb > 0 {
		for i:=1; i <= nb; i++ {
			result1 = result1 * i
	
		}
		return result1
	} else {
		return 0
	}

	
}