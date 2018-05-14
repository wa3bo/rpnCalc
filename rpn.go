package main

import (
	"fmt"
	"os"
)

/*func init() {
	handlers := [5]string{"+", "-", "/", "*", "exit"}
	stack := make([]string, 0)
}*/
var handlers = [5]string{"+", "-", "/", "*", "exit"}
var stack = make([]string, 0)

func main() {

	for {
		fmt.Printf(">> ")
		input := read()

		if input == "" {
			continue

		} else if input == "exit" {
			os.Exit(3)
		} else {
			checker := inArray(input, handlers)
			if checker {
				output(evaluate(input))
			} else {
				stack = append(stack, input)
			}
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

func output(data string) {
	fmt.Println(data)
}

func evaluate(data string) string {
	//do something
	return "hello"
}

func inArray(val string, array [5]string) (exists bool) {
	exists = false
	for _, v := range array {
		if val == v {
			exists = true
			return
		}
	}
	return
}
