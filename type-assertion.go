package main

import "fmt"

func random() interface{} {
	return 2
}

func main() {
	var res interface{} = random()
	// var resStr string = res.(string)
	// fmt.Println(resStr)

	// menggunakan switch
	switch val := res.(type) {
	case string:
		fmt.Println("value", val, "is string")
	case int:
		fmt.Println("value", val, "is integer")
	default:
		fmt.Println("value", val, "Unknown")
	}
}
