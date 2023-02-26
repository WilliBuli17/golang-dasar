package main

import "fmt"

/*
Type Assertions merupakan kemampuan merubah tipe data menjadi tipe data yang diinginkan
Fitur ini sering sekali digunakan ketika kita bertemu dengan data interface kosong
*/

func random() interface{} {
	return "OK"
}

func main() {
	result := random()

	resultString := result.(string)
	fmt.Println(resultString)

	resultInt := result.(int) // error karena returnnya di interface kosongnya string -> panic: interface conversion: interface {} is string, not int
	fmt.Println(resultInt)

	/*
		Saat salah menggunakan type assertions, maka bisa berakibat terjadi panic di aplikasi kita
		Jika panic dan tidak ter-recover, maka otomatis program kita akan mati
		Agar lebih aman, sebaiknya kita menggunakan switch expression untuk melakukan type assertions
	*/

	switch value := result.(type) {
	case string:
		fmt.Println("String,", value)
	case int:
		fmt.Println("Int,", value)
	default:
		fmt.Println("Unknown,", value)
	}
}
