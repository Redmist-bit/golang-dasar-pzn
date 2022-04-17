package main

import "fmt"

func main() {
	blacklist := func(name string) bool {
		return name == "anjing"
	}

	registerUser("bayu", blacklist)
	registerUser("anjing", blacklist)

	registerUser("root", func(name string) bool {
		return name == "root"
	})
	registerUser("bay", func(name string) bool {
		return name == "root"
	})
}

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You are blocked", name)
	} else {
		println("welcome", name)
	}
}
