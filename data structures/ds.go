package main

import "fmt"

func main() {
	var numbers [3]int

	numbers[0] = 1
	numbers[1] = 3
	numbers[2] = 98

	fmt.Println("Loop starts...")
	for i := 0; i < len(numbers); i++{
		fmt.Println(numbers[i])
	}

}