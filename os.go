package main
import (
	"os"
	"fmt"
)

func main(){
	args := os.Args
	fmt.Println(args)
	name, err := os.Hostname()
	if err == nil {
		fmt.Println(name)
	} else {
		fmt.Println("Error", err.Error())
	}

	username := os.Getenv("APP_USERNAME")
	password := os.Getenv("APP_PASSWORD")
	fmt.Println(username)
	fmt.Println(password)
}