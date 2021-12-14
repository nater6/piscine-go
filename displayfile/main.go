package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("File name missing")
	} else if len(args) > 1 {
		fmt.Println("Too many arguments")
	} else {

		file, err := os.ReadFile("quest8.txt")

		if err != nil {
			fmt.Print(err.Error())
		} else {
			fmt.Print(string(file))
		}
	}
}
