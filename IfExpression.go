package main

import "fmt"

func main() {
	name := "Ade"
	if name == "Bayu" {
		fmt.Println("Hai Bos", name)
	} else if name == "Ade" {
		fmt.Println("Hi Ade")
	} else {
		fmt.Println("HI Else")
	}

	if length := len(name); length > 5 {
		fmt.Println("Panjang")
	} else {
		fmt.Println("B aja")
	}
}
