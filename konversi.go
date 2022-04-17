package main

import "fmt"

func main() {
	var nilai32 int32 = 129
	var nilai64 int64 = int64(nilai32)
	var nilai8 int8 = int8(nilai32) // int8 gk cukup buat nampung karna kegedean variable awalnya

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai8)

	var name = "Bayu"
	var e byte = name[0] // byte == uint8
	var eString string = string(e)
	fmt.Println(name)
	fmt.Println(e)
	fmt.Println(eString)
}
