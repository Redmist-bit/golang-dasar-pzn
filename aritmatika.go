package main

import "fmt"

func main() {
	var a = 10
	var b = 7
	var c = a + b
	fmt.Println(c)

	// augmented assignments
	a *= 5 // a= 10 * 5
	fmt.Println(a)

	//unary operator
	a--
	fmt.Println(a)
}
