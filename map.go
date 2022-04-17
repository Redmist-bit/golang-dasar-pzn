package main

import "fmt"

func main() {
	person := map[string]string{
		"name":    "Bayu",
		"address": "Harapan Baru",
	}
	person["title"] = "Programmer"
	fmt.Println(person["name"])
	fmt.Println(person)

	fmt.Println(len(person))

	book := make(map[string]string)
	book["title"] = "Buku Go-Lang"
	book["author"] = "Bayu"
	fmt.Println(book)

	delete(book, "author")
	fmt.Println(book)
}
