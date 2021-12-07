package piscine

func IsPrime(nb int) bool {
	result1 := true
	if nb == 2 {
		result1 = true
	} else if nb > 2 {
		for i := 2; i < nb; i++ {
			if nb%i == 0 {
				return false
			}
		}
	} else if nb <= 1 {
		result1 = false
	}
	return result1
}
