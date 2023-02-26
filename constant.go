package main

import "fmt"

func main() {
	/*
	* Constant adalah variable yang nilainya tidak bisa diubah lagi setelah pertama kali diberi nilai
	* Cara pembuatan constant mirip dengan variable, yang membedakan hanya kata kunci yang digunakan adalah const, bukan var
	* Saat pembuatan constant, kita wajib langsung menginisialisasikan datanya
	 */

	const firstName string = "Willi"
	const lastName = "Buli" // bisa tanpa tipe data

	// error
	//firstName = "Tidak Dapat Diubah" //cannot assign to firstName (constant "Willi" of type string)
	//lastName = "Tidak Dapat Diubah" //cannot assign to lastName (untyped string constant "Buli")

	fmt.Println(firstName)
	fmt.Println(lastName)

	// Sama seperti variable, di Go-Lang  juga kita bisa membuat constant secara sekaligus banyak

	const (
		hobby        = "eat"
		favoriteFood = "fried chicken"
	)

	fmt.Println(hobby)
	fmt.Println(favoriteFood)
}
