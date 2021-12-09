package piscine

func Index(s string, toFind string) int {
	x := len(s)
	y := len(toFind)

	for i := 0; i < x-y; i++ {
		if toFind == s[i:i+y] {
			return i
		}
	}
	return -1
}
