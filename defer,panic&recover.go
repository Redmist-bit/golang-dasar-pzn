package main

import "fmt"

func main() {
	// runApp(0)
	runAppPanic(true)
	fmt.Println("jalan tetap")
}

func logging() {
	fmt.Println("Selesa memanggil logging")
}

// func runApp(val int) {
// 	defer logging() //defer. jalanin function walaupun error. ditaru di awal
// 	fmt.Println("Run App")
// 	res := 10 / val
// 	fmt.Println(res, "result")
// 	// logging()
// }

func endApp() {
	//recover untuk mengambil data panic
	message := recover()
	if message != nil {
		fmt.Println("Err message:", message)
	}
	fmt.Println("done")
}

func runAppPanic(error bool) {
	defer endApp()
	if error {
		panic("APK STOP")
	}
	fmt.Println("APK RUN")
}

