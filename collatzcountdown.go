package piscine

func CollatzCountdown(start int) int {
	if start <= 0 {
		return -1
	}

	x := start
	i := 0
	for x != 1 {
		if x%2 == 0 {
			x /= 2
		} else {
			x = (3 * x) + 1
		}
		i++
	}

	return i
}
