package piscine

func StrLen(s string) int {
	x := []rune(s)
	j := 0
	for i := 0; i < len(x); i++ {
		j++
	}
	return j
}