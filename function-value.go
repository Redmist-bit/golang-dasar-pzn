package main
import "fmt"
func main()  {
	goodbye:=getGoodBye
	result := goodbye("bay")
	fmt.Println(result)
}
func getGoodBye(name string) string {
	return "Good Bye" + name
}