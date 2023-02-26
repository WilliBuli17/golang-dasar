package main

import "fmt"

/*
Defer function adalah function yang bisa kita jadwalkan untuk dieksekusi setelah sebuah function selesai di eksekusi
Defer function akan selalu dieksekusi walaupun terjadi error di function yang dieksekusi
*/

func logging() {
	fmt.Println("Selesai memanggil function")
}

func runAplication(value int) {
	defer logging()
	fmt.Print("Run Applicarion")
	result := 10 / value
	fmt.Println("Result, ", result)
}

func main() {
	runAplication(10)
}
