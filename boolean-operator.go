package main

import "fmt"

func main() {
	var test = 80
	var absen = 90

	var lulusTes = test > 70
	var lulusAbsen = absen > 80
	fmt.Println(lulusTes)
	fmt.Println(lulusAbsen)

	var lulus = lulusTes && lulusAbsen
	fmt.Println(lulus)
}
