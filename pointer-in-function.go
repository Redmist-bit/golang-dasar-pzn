package main
import "fmt"

func changeAddress(addres *Address){
	addres.Country = "Indonesia"
}

type Address struct {
	City, Province, Country string
}

func main(){
	addres := &Address{"Samarinda","Kaltim",""}
	changeAddress(addres)
	fmt.Println(addres)
}