package piscine

func RecursiveFactorial(nb int) int {
	result1 := 0
	if nb == 0 && nb < 21{
		result1 = 1
	} else if nb > 0 {
		result1 = nb * RecursiveFactorial(nb -1)
	} else {
		return 0
	}
	return result1
}