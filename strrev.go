package piscine

func StrRev(s string) string {
	rune1 := []rune(s)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		x := rune1[i]
		rune1[i] = rune1[j]
		rune1[j] = x
	}
	y := string(rune1)
	return y
}
