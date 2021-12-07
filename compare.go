package piscine

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

	if i == j {
		return 0
	} else if i < j {
		return -1
	} else {
		return 1
	}
}
