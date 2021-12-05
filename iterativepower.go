package piscine

func IterativePower(nb int, power int) int {
	result1 := 0
	if power > 0 {
		for i := 0; i < power; i++ {
			result1 = nb * nb
		}
	} else if power == 0 {
		result1 = 1
	} else {
		return 0
	}
	return result1
}
