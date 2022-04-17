package main

import "fmt"

func main() {
	fmt.Println(getHello("Bay"))
	fmt.Println(getHello(""))

	first, last := getFullName()
	fmt.Println(first, last)
	satuAj, _ := getFullName()
	fmt.Println(satuAj)

	fmt.Println(getCompleteName())

	a, b, c := getCompleteName()
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}

func getHello(name string) string {
	if name == "" {
		return "Hello Bro"
	} else {
		return "Hello " + name
	}
}

func getFullName() (string, string) {
	return "Bayu", "Ade"
}

// named return values

func getCompleteName() (first, middle, last string) {
	first = "Mohamad"
	middle = "Bayu Ade"
	last = "Irawan"
	return
}
