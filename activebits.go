package piscine

func ActiveBits(n int) int {
	byteCheck := []int{128, 64, 32, 16, 8, 4, 2, 1}
	activeCheck := []int{}
	x := n

	for i := 0; i < len(byteCheck); i++ {
		answer := x / byteCheck[i]
		activeCheck = append(activeCheck, answer)
		x = x % byteCheck[i]
	}

	counter := 0

	for i := 0; i < len(activeCheck); i++ {
		if activeCheck[i] == 1 {
			counter++
		}
	}

	return counter
}
