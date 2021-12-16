package piscine

func Atoi(s string) int {
	r := s
	rune1 := s[1:]
	negative := false

	if len(s) == 0 {
		return 0
	}
	// Check all vakues are digits
	if r[0] == '-' || r[0] == '+' {
		for i := 1; i < len(s); i++ {
			if !(r[i] >= '0' && r[i] <= '9') {
				return 0
			}
		}
		if r[0] == '-' {
			negative = true
		}
	} else {
		for i := 0; i < len(s); i++ {
			if !(r[i] >= '0' && r[i] <= '9') {
				return 0
			}
		}
	}

	// Turn the string into a digit
	myDigit := 0
	if r[0] == '-' || r[0] == '+' {
		for i := 0; i < len(rune1); i++ {
			myDigit = myDigit * 10
			myDigit = myDigit + int(rune1[i]-48)
		}
	} else {
		for _, number := range r {
			myDigit = myDigit * 10
			myDigit = myDigit + int(number-48)
		}
	}

	// Turn the number negative
	if negative {
		myDigit *= -1
	}

	return myDigit
}
