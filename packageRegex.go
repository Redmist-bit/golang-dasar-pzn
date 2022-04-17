package main
import (
	"regexp"
	"fmt"
)

func main(){
	var regex *regexp.Regexp = regexp.MustCompile(`a([a-z])e`)
	fmt.Println(regex.MatchString("ade"))
	fmt.Println(regex.MatchString("ace"))
	fmt.Println(regex.MatchString("aDa"))
}