package main

import "fmt"

func main()  {
	var bay Customer
	bay.Nama = "Bayu Ade"
	bay.Alamat = "HarapanBaru"
	bay.Umur = 24

	bay.sayHi("andre")
}

func (customer Customer) sayHi(name string)  {
	fmt.Println("Hai Bos", name, "My Name is", customer.Nama)
}

// struct representasi dari data
type Customer struct {
	Nama, Alamat string
	Umur int
}
