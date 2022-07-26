package piscine

func TrimAtoi(s string) int {
	r := []rune(s)
	intValue := 0
	check := false
	noDig := true

	for i := 0; i < len(s); i++ {
		if r[i] >= '0' && r[i] <= '9' {
			rDigit := r[i] - 48
			intValue = intValue * 10
			intValue = intValue + int(rDigit)
			noDig = false
		} else if r[i] == '-' && intValue == 0 {
			check = true
		} 
	}

	if noDig {
		return 0
	} else if check {
		return intValue * -1
	} else {
		return intValue
	}
}
