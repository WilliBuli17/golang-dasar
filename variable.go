package main

import "fmt"

func main() {
	/*
	* Variable adalah tempat untuk menyimpan data
	* Variable digunakan agar kita bisa mengakses data yang sama dimanapun kita mau
	* Di Go-Lang Variable hanya bisa menyimpan tipe data yang sama, jika kita ingin menyimpan data yang berbeda-beda jenis, kita harus membuat beberapa variable
	* Untuk membuat variable, kita bisa menggunakan kata kunci var, lalu diikuti dengan nama variable dan tipe datanya
	 */

	var name string

	name = "Styephen William Buli"
	fmt.Println(name)

	name = "Willi Buli"
	fmt.Println(name)

	/*
	* Saat kita membuat variable, maka kita wajib menyebutkan tipe data variable tersebut
	* Namun jika kita langsung menginisialisasikan data pada variable nya, maka kita tidak wajib menyebutkan tipe data variable nya
	 */

	var hobby = "eat"
	fmt.Println(hobby)

	hobby = "sleep"
	fmt.Println(hobby)

	/*
	* Di Go-Lang, kata kunci var saat membuat variable tidak lah wajib.
	* Asalkan saat membuat variable kita langsung menginisialisasi datanya
	* Agar tidak perlu menggunakan kata kunci var, kita perlu menggunakan kata kunci := saat menginisialisasikan data pada variable tersebut
	 */

	favoriteFood := "fried chicken"
	fmt.Println(favoriteFood)

	favoriteFood = "fried rice" // yang keduandan seterusnya ridak perlu := tapi hanya = saja
	println(favoriteFood)

	/*
	* Di Go-Lang kita bisa membuat variable secara sekaligus banyak
	* Code yang dibuat akan lebih bagus dan mudah dibaca
	 */

	var (
		firstName = "Willi"
		lastname  = "Buli"
	)

	fmt.Println(firstName)
	fmt.Println(lastname)
}
