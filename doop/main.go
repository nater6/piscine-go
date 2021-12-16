package main

import (
	"os"
	"strconv"
)

func main() {
	arguements := os.Args
	args := arguements[1:]
	// Check the length is the same
	if len(args) != 3 {
		return
	}
	// Check 1st arg is an int
	for _, digit := range args[0] {
		if !(digit >= '0' && digit <= '9') {
			return
		}
	}
	// Check the 3rd arg is an int
	for _, digit := range args[2] {
		if !(digit >= '0' && digit <= '9') {
			return
		}
	}
	// Check that the middle arg is an operator
	if !(args[1] == "+" || args[1] == "-" || args[1] == "/" || args[1] == "%" || args[1] == "*") {
		return
	}
	// Check that modulo isnt with 0
	if args[1] == "%" && args[2] == "0" {
		print("No modulo by 0")
	}
	// Check division isnt with a zero
	if args[1] == "/" && args[2] == "0" {
		print("No division by 0")
	}
	// Check number is within int64 range

	int1, _ := strconv.Atoi(args[0])
	int2, _ := strconv.Atoi(args[2])
	if int1 >= 9223372036854775807 || int1 < -9223372036854775808 {
		return
	}
	if int2 >= 9223372036854775807 || int1 < -9223372036854775808 {
		return
	}
	result := 0

	if args[1] == "+" {
		result = int1 + int2
	} else if args[1] == "-" {
		result = int1 - int2
	} else if args[1] == "*" {
		result = int1 * int2
	} else if args[1] == "/" {
		result = int1 / int2
	} else if args[1] == "%" {
		result = int1 % int2
	}
	print(result)
}
