package piscine

func BasicAtoi(s string) int {
	rune1 := []rune(s)
	for i := 0; i<len(s); i++ {
		if (48 <= rune1[i] && 57<= rune1[i]){
			x := int(rune1[i])
		} else {
			return 0
		}
		


	}
}
