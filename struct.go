package main

import "fmt"

func main()  {
	var bay Customer
	bay.Nama = "Bayu Ade"
	bay.Alamat = "HarapanBaru"
	bay.Umur = 24

	fmt.Println(bay)
	fmt.Println(bay.Alamat)
	// simple struct
	rere := Customer{
		Nama: "Rereh Nur Wening",
		Alamat: "Jateng",
		Umur: 24,
	}
	fmt.Println(rere)
	nuning:= Customer{"Nuning","Kndl",24}
	fmt.Println(nuning)
}

// struct representasi dari data
type Customer struct {
	Nama, Alamat string
	Umur int
}
