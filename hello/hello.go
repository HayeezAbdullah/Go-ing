package main

import "fmt"

func main() {
	const Pi = 3.14159
	var r int = 34
	var another_r int = 66

	fmt.Println("r is bigger than another_r ? -> ", (2 * Pi * float64(r)) > (2 * Pi * float64(another_r)))

}