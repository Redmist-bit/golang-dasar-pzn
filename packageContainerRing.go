package main
import (
	"container/ring"
	"strconv"
	"fmt"
)

func main(){
	var data *ring.Ring = ring.New(5)
	for i := 0; i < data.Len(); i ++ {
		data.Value = "data "+ strconv.FormatInt(int64(i),10)
		data = data.Next()
	}
	data.Do(func(value interface{}) {
		fmt.Println(value)
	})
}