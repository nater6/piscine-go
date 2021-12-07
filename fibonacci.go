package piscine

func Fibonacci(index int) int {
	if index >= 2 {

		result := Fibonacci(index-2) + Fibonacci(index-1)
		return result

	} else if index == 0 || index == 1 {
		return index
	} else {
		return -1
	}
}
