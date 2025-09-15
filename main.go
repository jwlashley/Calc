package main

import (
	"fmt"
	"os"
	"strconv"
)

func displayHelp() {
	fmt.Print(`------------------------------------------------------------------------------------------------------
- Calc is an efficiently small calculator program to allow for simple arithmetic from the terminal.
- To utilize the program please use the following format:
--- calc <number1> <operator> <number2>
-- The following operators are available in this version: (+,-,x,*,/)
------------------------------------------------------------------------------------------------------
`)
}

func main() {

	// Get Args from the OS
	args := os.Args[1:]

	if len(args) < 3 {
		fmt.Println("Error: Invalid command usage.")
		fmt.Println("Usage: calc <number1> <operator> <number2>")
		if args[0] == "-help" {
			displayHelp()
		}
		os.Exit(1)
	}

	num1, err1 := strconv.ParseFloat(args[0], 64)
	if err1 != nil {
		fmt.Println("Error: The first number entered is invalid.")
		os.Exit(1)
	}

	num2, err2 := strconv.ParseFloat(args[2], 64)
	if err2 != nil {
		fmt.Printf("Error: The second number entered is invalid.")
		os.Exit(1)
	}

	var result float64

	operator := args[1]

	switch operator {
	case "+":
		result = num1 + num2

	case "-":
		result = num1 - num2
	case "x", "*":
		result = num1 * num2
	case "/":
		result = num1 / num2
	default:
		fmt.Println("Error: Invalid operator entered. Please use '+, -, x, or /'")
		os.Exit(1)
	}

	fmt.Printf("Result: %.2f \n", result)

}
