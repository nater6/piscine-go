package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(os.Args) == 0 {
		fmt.Println("File name missing")
	}

	if len(args) > 1 {
		fmt.Println("Too many arguments")
	}

	file, err := os.ReadFile("quest8.txt")

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(string(file))
	}
}
