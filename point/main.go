package main

import "fmt"

type pointers struct {
	x int
	y int
}

func main() {
	points := pointers{}
	points.x = 42
	points.y = 21

	fmt.Printf("x = %d, y = %d\n", points.x, points.y)
}
