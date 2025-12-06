package main

import "fmt"

func get_quantity(quantity int, availableStock int) ( string, error){
	if quantity <= 0 {
		return "Quantity must be positive", nil
	}
	if quantity > availableStock{
		return "Quantity not enough", nil
	}

	return "Order accepted", nil
}

func main() {
	var customerReq, inStock int

	fmt.Println("Enter customer requirement: ")
	fmt.Scanln(&customerReq)

	fmt.Println("Enter Instock: ")
	fmt.Scanln(&inStock)

	result, err := get_quantity(customerReq, inStock)
	if err != nil {
		fmt.Println("error instock", err)
		return
	}
	fmt.Println(result)
}