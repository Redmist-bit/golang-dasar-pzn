package main

import "fmt"

func main() {
	name := "Ae"

	switch name {
	case "Ade":
		fmt.Println("Halo, Ade")
	case "Bayu":
		fmt.Println("Halo, Bos")
	default:
		fmt.Println("DEFAULT")
	}

	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Panjang")
	case false:
		fmt.Println("B aj")
	}

	pjg := len(name)

	switch {
	case pjg > 10:
		fmt.Println("Lebih dari 10")
	case pjg < 5:
		fmt.Println("Kurang dari 5")
	}
}
