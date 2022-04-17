package main

import "fmt"

func main() {
	var names[4]string
	names[0] = "Mohamad"
	names[1] = "Bayu"
	names[2] = "Ade"
	names[3] = "Irawan"

	var data = [3]int{
		12,
		7,
		99,
	}
	fmt.Println(len(data))
	fmt.Println(data)
	fmt.Println(names[1])
}