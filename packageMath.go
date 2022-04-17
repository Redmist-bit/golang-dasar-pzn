package main
import (
	"math"
	"fmt"
)

func main(){
	fmt.Println(math.Round(1.7))
	fmt.Println(math.Round(1.3))
	fmt.Println(math.Floor(1.3))
	fmt.Println(math.Ceil(1.3))
	fmt.Println(math.Max(1.3,1.7))
	fmt.Println(math.Min(1.3,1.7))
}