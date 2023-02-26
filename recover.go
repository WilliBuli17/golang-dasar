package main

/*
Recover adalah function yang bisa kita gunakan untuk menangkap data panic
Dengan recover proses panic akan terhenti, sehingga program akan tetap berjalan
*/

import "fmt"

func endAppRecover() {
	message := recover() // di sarankan ada dalam defer atau di luar fungsi yang ada panic
	if message != nil {
		fmt.Println("Ther's an error, ", message)
	}
	fmt.Println("Ended App")
}

func runRecover(error bool) {
	defer endAppRecover()
	if error {
		panic("App Error")
	}

	fmt.Println("App Runing")
}
func main() {
	runRecover(false)
}
