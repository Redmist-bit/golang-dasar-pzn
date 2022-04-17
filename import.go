package main
import "fmt"
import "belajar-golang-dasar/helper"
func main() {
	helper.SayHello("Bay")
	fmt.Println(helper.App)
	// access modifier - tidak bisa akses function package sayGoodBye awalan huruf kecil
	// helper.sayGoodBye("Bay")
	// fmt.Println(helper.ver)
}
