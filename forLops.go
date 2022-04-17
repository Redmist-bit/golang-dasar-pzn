package main

import "fmt"

func main() {
	count := 1

	for i := 10; i >= count; {
		fmt.Println("ke:", count)
		count++
	}

	// for count <= 10 {
	// 	fmt.Println("Loop ke:", count)
	// 	count++
	// }
	for i := 1; i < 10; i++ {
		fmt.Println("ke:", i)
	}

	slice := []string{"bayu", "ade"}
	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	for index, val := range slice {
		fmt.Println("index Ke-", index, "=", val)
	}
	// kalo mau ambil value doang, biar gak eror indexnya di ganti underscore _
	for _, val := range slice {
		fmt.Println(val)
	}

	book := make(map[string]string)
	book["title"] = "Buku Go-Lang"
	book["author"] = "Bayu"

	for key, val := range book {
		fmt.Println("Key -", key, "=", val)
	}
}
