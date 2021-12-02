package piscine

func StrLen(s string) int {
	j := 0
	for i := 0; i < len(s); i++ {
		j++
	}
	return j
}
