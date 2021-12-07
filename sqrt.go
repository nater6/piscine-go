package piscine

func Sqrt(nb int) int {
	var result int

	if nb > 0 {
		for x := 1; x < nb+1; x++ {
			if nb == x*x {
				result = x
			}
		}
	} else {
		result = 0
	}
	return result
}
