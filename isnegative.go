package piscine

import "fmt"

func IsNegative(nb int) {
	if nb < 0 {
		fmt.Println('T')
		fmt.Println('\n')
	} else {
		fmt.Println('F')
		fmt.Println('\n')
	}
}
