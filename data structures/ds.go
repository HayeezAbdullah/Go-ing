package main

import "fmt"

func main() {
	superPizza := []string {"cheese", "chicken", "veggie","bbq", "mutton", "beef", "prawns"}
	normalPizza := superPizza[0:4]

	fmt.Println("BEFORE CHANGE")
	fmt.Println("Super Pizza: ", superPizza)
	fmt.Println("Normal Pizza: ", normalPizza)
	fmt.Println(".........................")

	normalPizza = append(normalPizza, "tandoori", "masala")

	fmt.Println("AFTER CHANGE")
	fmt.Println("Super Pizza: ", superPizza)
	fmt.Println("Normal Pizza: ", normalPizza)
	fmt.Println(".........................")

	
	
}