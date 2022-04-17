package main

import "fmt"

func main() {
	sayHi()
	sayHi()
	for i := 0; i < 3; i++ {
		sayHi()
	}
	sayWithParam("Bay", "Ade")
}

func sayHi() {
	fmt.Println("TEST FUNCTION")
}

func sayWithParam(firstName string, lastName string) {
	fmt.Println("Hi", firstName, lastName)
}
