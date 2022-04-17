package main
// reference from array
import "fmt"

func main() {
	var months = [...]string{
		"januari",
		"februari",
		"maret",
		"april",
		"mei","juni","Juli","Agustuts","September","Oktober","November","Desember",
	}
	var slice1 = months[2:5]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	months[3] = "DiUbah"
	fmt.Println(slice1)
	slice1[0] = "diGanti"
	fmt.Println(slice1)
	fmt.Println(months)

	var slice2 = months[10:]
	fmt.Println(slice2)
	var slice3 = append(slice2,"Bayu") // nambah baru tp udh full, jadi bikin array baru
	fmt.Println(slice3)
	slice3[1] = "DesemberGanti"
	fmt.Println(slice3)
	fmt.Println(slice2)
	fmt.Println(months)

	newSlice := make([]string,2,5)
	newSlice[0] = "Bayu"
	newSlice[1] = "Ade"
	// newSlice[2] = "irawan"
	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice,newSlice)
	fmt.Println("copyslice",copySlice)

	iniArray := [5]int{1,2,3,4,5}
	iniArrayJuga := [...]int{1,2,3,4,5}
	iniSlice := []int{1,2,3,4,5}

	fmt.Println("iniArray",iniArray)
	fmt.Println("iniArrayJuga",iniArrayJuga)
	fmt.Println("iniSlice",iniSlice)
}
