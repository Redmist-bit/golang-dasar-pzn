package main
import (
	"sort"
	"fmt"
)

type User struct {
	Name string
	Age int
}

type UserSlice []User

func (userSlice UserSlice) Len() int {
	return len(userSlice)
}

func (value UserSlice) Less(i,j int) bool {
	return value[i].Age < value[j].Age
}

func (value UserSlice) Swap(i,j int) {
	value[i], value[j] = value[j], value[i]
}
func main(){
	users := []User {
		{"Bayu",24},
		{"Ade",30},
		{"Irawan",25},
	}

	fmt.Println(users)

	sort.Sort(UserSlice(users))

	fmt.Println(users)
}