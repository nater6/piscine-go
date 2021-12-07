package piscine

func RecursivePower(nb int, power int) int {
	if power > 1 {
		result := RecursivePower(nb, power-1) * nb
		return result
	} else if power == 1 {
		return nb
	} else if power == 0 {
		return 1
	} else {
		return 0
	}
}
