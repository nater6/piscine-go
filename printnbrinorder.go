package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	var numbers []int

	if n < 0 {
		return
	} else if n == 0 {
		numbers = append(numbers, 0)
	}

	for n > 0 {
		numbers = append(numbers, n%10)
		n = n / 10
	}

	for i := 0; i < len(numbers)-1; i++ {
		for j := 0; j < len(numbers)-1; j++ {
			if numbers[j] > numbers[j+1] {
			numbers[j], numbers[j+1] = numbers[j+1], numbers[j]
			}
	}
		}

	for _, digit := range numbers {
		z01.PrintRune(rune(digit) + 48)
	}
}
