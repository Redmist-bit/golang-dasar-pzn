package main
import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married(){
	man.Name = "Mr. " + man.Name
}
func main(){
	bay := Man{"Bay"}
	bay.Married()
	fmt.Println(bay)
}