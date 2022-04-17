package main
import (
	"strconv"
	"fmt"
)

func main(){
	boolean, err := strconv.ParseBool("true")// convert string true menjadi boolean true

	if err == nil {
		fmt.Println(boolean)
	} else {
		fmt.Println(err.Error())
	}
}