package main

import "fmt"

func main()  {
	var bay Person
	bay.Name = "BAYU"
	sayHello(bay)
}

func sayHello(hasName HasName)  {
	fmt.Println("Hai Bos", hasName.GetName())
}

type HasName interface {
	GetName() string
}

type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}


