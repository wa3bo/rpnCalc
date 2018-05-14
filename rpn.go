package main

import (
	"fmt"
	"os"
	"strconv"
)

type stack struct {
	vec []int
}

var handlers = [5]string{"+", "-", "/", "*", "exit"}
var data = new(stack)

func main() {
	var checker bool

	for {
		fmt.Printf(">> ")
		input := read()

		if input == "" {
			continue

		} else if input == "exit" {
			os.Exit(3)
		} else {
			checker = inArray(input, handlers)
			if checker {
				output(evaluate(input))
			} else {
				checker = isCharacter(input)
				if checker {
					insert, err := strconv.Atoi(input)
					if err != nil {
						fmt.Println("fml")
					}
					data.put(insert)
				} else {
					fmt.Println("syntax error")
					continue
				}

			}
		}

	}

}

//read inputs
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

//handle operations
func evaluate(z string) string {
	if data.empty() {
		return "stack is empty"
	} else {
		if z == "+" {
			data.put(data.pop() + data.pop())
		} else if z == "/" {
			erg := data.pop()
			erg2 := data.pop()
			if erg2 == 0 {
				fmt.Println("div by 0")
				os.Exit(3)
			} else {
				data.put(erg2 / erg)
			}
		} else if z == "*" {
			data.put(data.pop() * data.pop())
		} else if z == "-" {
			data.put(data.pop() - data.pop())
		}
	}
	fmt.Println(data.vec)
	return strconv.Itoa(data.vec[0])
}

//check if a operation is to do
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

//check if it is a digit
func isCharacter(val string) (erg bool) {
	for i := 0; i < len(val); i++ {
		erg = false
		if val[i] >= 48 && val[i] <= 57 {
			erg = true
		}
	}
	return
}

//for stack operations
func (s stack) empty() bool {
	return len(s.vec) == 0
}
func (s *stack) put(i int) {
	s.vec = append(s.vec, i)
}
func (s *stack) pop() int {
	d := s.vec[len(s.vec)-1]
	s.vec = s.vec[:len(s.vec)-1]
	return d
}
