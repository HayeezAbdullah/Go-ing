package main

import "fmt"

func main() {
	
var age int
var count int
	for{

	fmt.Println("Enter age: ")
	fmt.Scanln(&age)
	checkAge(age)
	count = count + 1
	if count == 5 {
		fmt.Println("House Full")
		break
	}
	}
}

func checkAge(age int){
	switch {
	case age < 18:
		fmt.Println("Minor")
	case (age > 17) && (age < 59):
		fmt.Println("Adult")
	default:
		fmt.Println("Old")	
	}
}