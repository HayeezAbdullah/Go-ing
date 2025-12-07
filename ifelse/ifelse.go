package main

import "fmt"

func main() {
	var age int
	fmt.Println("Enter age: ")
	fmt.Scanln(&age)

	switch {
	case age < 18:
		fmt.Println("Minor")
	case (age > 17) && (age < 59):
		fmt.Println("Adult")
	default:
		fmt.Println("Old")	
	}

}