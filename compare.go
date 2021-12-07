package piscin

func Compare(a, b string) int {
	compareA := []rune(a)
	compareB := []rune(b)

	i := 0
	for i < len(compareA) {
		i++
	}

	j := 0
	for j < len(compareB) {
		j++
	}
	result := 5
	if i == j {
		result = 0
	} else if i < j {
		result = (1)
	} else {
		result = -1
	}
	return result
}
