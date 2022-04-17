package main

import "fmt"

type Address struct {
	City, Province, Country string
}
func main(){
	// pass by value
	// addrs1 := Address{"Samarinda","Kaltim","indonesia"}
	// addrs2 := addrs1
	// addrs2.City = "Semarang"

	// fmt.Println(addrs1)
	// fmt.Println(addrs2)
	//

	// pass by reference
	// addrs1 := Address{"Samarinda","Kaltim","indonesia"}
	// addrs2 := &addrs1
	// addrs2.City = "Semarang"

	// fmt.Println(addrs1)
	// fmt.Println(addrs2)
	
	// Melakukan perubahan tanpa merubah addrs1
	// addrs1 := Address{"Samarinda","Kaltim","indonesia"}
	// addrs2 := &addrs1
	// addrs2.City = "Semarang"

	// addrs2 = &Address{"tgr","timur","idn"}

	// fmt.Println(addrs1)
	// fmt.Println(addrs2)

	// Melakukan perubahan semua
	// addrs1 := Address{"Samarinda","Kaltim","indonesia"}
	// addrs2 := &addrs1
	// addrs2.City = "Semarang"

	// *addrs2 = Address{"tgr","timur","idn"}

	// fmt.Println(addrs1)
	// fmt.Println(addrs2)

	// membuat baru dr kosong / new
	var addrs4 *Address = new(Address)
	fmt.Println(addrs4)
	addrs4.City = "Kendal"
	fmt.Println(addrs4)
}