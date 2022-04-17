package main
import (
	"errors"
	"fmt"
)

func Pembagi(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Pembagi tidak boleh 0")
	} else {
		result := nilai/pembagi
		return result, nil
	}
}

func main() {
	var contohErr error = errors.New("Ups error")
	fmt.Println(contohErr.Error())
	hasil, err := Pembagi(100,0)
	if err == nil {
		fmt.Println("Hasil",hasil)
	} else {
		fmt.Println("Error",err.Error())
	}
}
