package main

import "fmt"

func calculate(num1, num2 float64, operator string) (float64, error) {
	switch operator {
	case "+":
		return num1 + num2, nil
	case "-":
		return num1 - num2, nil
	case "*":
		return num1 * num2, nil
	case "/":
		if num2 == 0 {
			return 0, fmt.Errorf("cannot divide by zero")
		}
		return num1 / num2 , nil
	default:
		return 0, fmt.Errorf("invalid input")
	}
}

func main() {
	var input1 , input2 float64
	var operator_string string

	fmt.Println("Enter first number: ")
	fmt.Scanln(&input1)
	fmt.Println("Enter operator: ")
	fmt.Scanln(&operator_string)
	fmt.Println("Enter second number: ")
	fmt.Scanln(&input2)

	result, err := calculate(input1, input2, operator_string)
	if err != nil {
		fmt.Println("Caught error: ", err)
		return
	}
	fmt.Printf(`Result is %v`, result)
}