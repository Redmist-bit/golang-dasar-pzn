package main
import (
	"strings"
	"fmt"
)

func main(){
	fmt.Println(strings.Contains("Bayu Ade","Ade"))
	fmt.Println(strings.Contains("Bayu Ade","irawan"))
	fmt.Println(strings.Split("Bayu Ade"," "))
	fmt.Println(strings.ToLower("Bayu Ade"))
	fmt.Println(strings.ToUpper("Bayu Ade"))
	fmt.Println(strings.ToTitle("Bayu Ade"))
	fmt.Println(strings.Trim("   Bayu Ade    "," "))
	fmt.Println(strings.ReplaceAll("Bayu Ade Bayu","Bayu","Irawan"))
}