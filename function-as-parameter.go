package main

import "fmt"

func main() {
	sayHelloWithFilter("Bayu", spamFilter)
	sayHelloWithFilter("Anjing", spamFilter)
	// filter :=
}

//with typ declaration
type Filter func(string) string

func sayHelloWithFilter(name string, filter Filter) {
	nameFiltered := filter(name)
	fmt.Println("Helo ", nameFiltered)
}

// func sayHelloWithFilter(name string, filter func (string) string) {
// 	nameFiltered := filter(name)
// 	fmt.Println("Helo ", nameFiltered)
// }

func spamFilter(name string) string {
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}
}
