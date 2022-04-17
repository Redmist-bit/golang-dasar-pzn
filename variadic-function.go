package main

import "fmt"

func sumAll(numbers ...int) int {
	total := 0
	for _, v := range numbers {
		total += v
	}
	return total
}

func main() {
	total := sumAll(10, 20)
	fmt.Println(total)

	slice := []int{10, 20, 30}
	total = sumAll(slice...)
	fmt.Println(total)
}
