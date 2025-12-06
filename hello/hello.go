package main

import (
	"fmt"
)

func main() {
	var num1 float64
	var num2 float64
	var operator string

	fmt.Println("Enter the first number: ")
	fmt.Scanln(&num1)
	fmt.Println("Enter the second number: ")
	fmt.Scanln(&num2)
	fmt.Println("Enter Operator (+,-,*,/): ")
	fmt.Scanln(&operator)
	
	switch operator {
	case "+":
		fmt.Printf(`Result %v`, num1 + num2)
	case "-":
		fmt.Printf(`Result %v`, num1 - num2)
	case "*":
		fmt.Printf(`Result %v`, num1 * num2)
	case "/":
		fmt.Printf(`Result %v`, num1 / num2)
	default:
		fmt.Println("Invalid Operator")	
	}
}