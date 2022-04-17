package main

import "fmt"

func main() {
	type NoKtp string
	type Married bool
	var NoKtpAing NoKtp = "6472000"
	var Status Married = false
	fmt.Println(NoKtpAing)
	fmt.Println(Status)
}
