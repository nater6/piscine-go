package piscine

import "fmt"

func QuadE(x int, y int) {
	// Checking input values are positive
	if x > 0 && y > 0 {
		if y > 1 {
			// Top horizontal line of rectangle
			if x == 1 {
				fmt.Println("A")
			} else {
				fmt.Print("A")
			}
			for a := x - 2; a > 0; a-- {
				fmt.Print("B")
			}
			if x > 1 {
				fmt.Println("C")
			}

			// Middle verticle lines of rectangle

			for b := y - 2; b > 0; b-- {
				if x == 1 {
					fmt.Println("B")
				} else {
					fmt.Print("B")
					for a := x - 2; a > 0; a-- {
						fmt.Print(" ")
					}
				}
				if y > 1 && x != 1 {
					fmt.Println("B")
				}
			}
			// Bottom line of rectangle
			if x == 1 {
				fmt.Println("C")
			} else {
				fmt.Print("C")
			}
			for a := x - 2; a > 0; a-- {
				fmt.Print("B")
			}
			if x > 1 {
				fmt.Println("A")
			}
		} else {
			// Top horizontal line of rectangle
			if x == 1 {
				fmt.Println("A")
			} else {
				fmt.Print("A")
				for a := x - 2; a > 0; a-- {
					fmt.Print("B")
				}
				if x > 1 {
					fmt.Println("C")
				}
			}
		}
	} else {
	}
}
