package piscine

func IterativePower(nb int, power int) int {
	result1 := 0
	x := nb
	if power > 1 {
		for i := 0; i < power-1; i++ {
			x = x * nb
			result1 = x
		}
		return result1
	} else if power == 1 {
		return nb
	} else if power == 0 {
		result1 = 1
		return result1
	} else {
		return 0
	}
}
