package main

import (
	"fmt"
	"os"
)

/*func init() {
	handlers := [5]string{"+", "-", "/", "*", "exit"}
	stack := make([]string, 0)
}*/

func main() {
	handlers := [5]string{"+", "-", "/", "*", "exit"}
	stack := make([]string, 0)

	for {
		fmt.Printf(">> ")
		input := read()

		if input == "" {
			continue

		} else if input == "exit" {
			os.Exit(3)
		} else {

			fmt.Println(input)
			stack = append(stack, input)
			fmt.Println(stack)
			fmt.Println(handlers[1])
		}

	}

}

func read() string {
	var number string
	fmt.Scanln(&number) //aber
	if number == "exit" {
		return "exit"
	} else if number == "" {
		return ""
	} else {
		return number
	}
}
