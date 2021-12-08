package piscine 

func TrimAtoi(s string) int {
	r := []rune(s)
	intValue :=0
	check := false

	for x:=0; (r[x]>='0'&&r[x]<='9')|| x <len(r); x++ {
		if r[x] == '-' {
			check = true
		}
	}

	for i:=0; i<len(s); i++ {
		if (r[i]>='0'&&r[i]<='9') {
		for i:=0; i<len(s); i++ {
			if (r[i]>='0'&&r[i]<='9') {
			rDigit := r[i] - 48
			intValue = intValue * 10
			intValue = intValue + int(rDigit)
		}

		}
		
		} else {
			return 0
		}
	}
	if check {
		return intValue * -1
	} else {
		return intValue
	}
}

