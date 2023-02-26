package main

import "fmt"

/*
Panic function adalah function yang bisa kita gunakan untuk menghentikan program
Panic function biasanya dipanggil ketika terjadi error pada saat program kita berjalan
Saat panic function dipanggil, program akan terhenti, namun defer function tetap akan dieksekusi
*/

func endApp() {
	fmt.Println("Ended App")
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("App Error")
	}
	fmt.Println("App Runing")
}
func main() {
	runApp(true)
}
