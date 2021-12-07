package piscine

func Sqrt(nb int) int {
	var result int

	if nb > 0 {
		for x:=0; x < nb; x++ {
			if nb == x*x {
				result = x
			} 
		} 
	} else {
		result = 0
	} 
	return result
}