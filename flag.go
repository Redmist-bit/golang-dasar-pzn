package main
import (
	"flag"
	"fmt"
)

func main(){
	var host *string = flag.String("host", "localhost", "Put Your Host")
	var user *string = flag.String("user", "root", "Put Your user")
	var password *string = flag.String("password", "root", "Put Your password")

	flag.Parse()

	fmt.Println("Host :",*host)
	fmt.Println("User :",*user)
	fmt.Println("Password :",*password)
}