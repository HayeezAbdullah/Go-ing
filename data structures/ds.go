package main

import "fmt"

func main() {
	pizza := []string {"cheese", "chicken", "veggie"}

	fmt.Println("length: ", len(pizza))
	fmt.Println("capacity: ", cap(pizza))

	pizza = append(pizza, "bbq")

	fmt.Println(pizza)
	fmt.Println("length: ", len(pizza))
	fmt.Println("capacity: ", cap(pizza))
	pizza = append(pizza, "beef")
	pizza = append(pizza, "mutton")
	pizza = append(pizza, "prawns")

	fmt.Println(pizza)
	fmt.Println("length: ", len(pizza))
	fmt.Println("capacity: ", cap(pizza))
}