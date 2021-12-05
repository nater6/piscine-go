package piscine

import "fmt"

func QuadB(x int, y int) {
	// Checking input values are positive
	if x > 0 && y > 0 {
		if y > 1 {
			// Top horizontal line of rectangle
			if x == 1 {
				fmt.Println("/")
			} else {
				fmt.Print("/")
			}
			for a := x - 2; a > 0; a-- {
				fmt.Print("*")
			}
			if x > 1 {
				fmt.Println("\\")
			}

			// Middle verticle lines of rectangle

			for b := y - 2; b > 0; b-- {
				if x == 1 {
					fmt.Println("*")
				} else {
					fmt.Print("*")
					for a := x - 2; a > 0; a-- {
						fmt.Print(" ")
					}
				}
				if y > 1 && x != 1 {
					fmt.Println("*")
				}
			}
			// Bottom line of rectangle
			if x == 1 {
				fmt.Println("\\")
			} else {
				fmt.Print("\\")
			}
			for a := x - 2; a > 0; a-- {
				fmt.Print("*")
			}
			if x > 1 {
				fmt.Println("/")
			}
		} else {
			// Top horizontal line of rectangle
			if x == 1 {
				fmt.Println("/")
			} else {
				fmt.Print("/")
				for a := x - 2; a > 0; a-- {
					fmt.Print("*")
				}
				if x > 1 {
					fmt.Println("\\")
				}
			}
		}
	} else {
	}
}
