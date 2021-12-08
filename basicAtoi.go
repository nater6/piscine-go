package piscine

func BasicAtoi(s string) int {
	x:=0
	r := []rune(s)

	for i:=0; i<len(s); i++ {
		y := r[i]-48
		x = x*10
		x += int(y)
	}
	return x
}

	

