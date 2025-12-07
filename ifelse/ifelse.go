package main

import "fmt"

func main() {
	var age int = 12
	if age >= 18 {
		fmt.Println("Adult")
	} else {
		fmt.Println("Minor")
	}
}