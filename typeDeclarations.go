package main

import "fmt"

func main() {
	/*
	* Type Declarations adalah kemampuan membuat ulang tipe data baru dari tipe data yang sudah ada
	* Type Declarations biasanya digunakan untuk membuat alias terhadap tipe data yang sudah ada, dengan tujuan agar lebih mudah dimengerti
	 */

	type NoKTP string
	type Married bool

	var noKtpUser NoKTP = "4237158971340768"
	var isMarried Married = false

	fmt.Println(noKtpUser)
	fmt.Println(isMarried)
}
