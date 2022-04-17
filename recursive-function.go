package main
import "fmt"

func factorialloop(value int) int {
	res := 1
	for i := value; i > 0; i-- {
		res *= i
	}
	return res
}
func main()  {
	loop:=factorialloop(5)
	fmt.Println(loop)
	fmt.Println(5*4*3*2*1)
	fmt.Println("pake function", recursiveFactorial(5))
}

//recursive function
func recursiveFactorial(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * recursiveFactorial(value-1)
	}
}