package main

import "fmt"

/*
Saat kita membuat parameter di function, secara default adalah pass by value, artinya data akan di copy lalu dikirim ke function tersebut
Oleh karena itu, jika kita mengubah data di dalam function, data yang aslinya tidak akan pernah berubah.
Hal ini membuat variable menjadi aman, karena tidak akan bisa diubah
Namun kadang kita ingin membuat function yang bisa mengubah data asli parameter tersebut
Untuk melakukan ini, kita juga bisa menggunakan pointer di function
Untuk menjadikan sebuah parameter sebagai pointer, kita bisa menggunakan operator * di parameternya
*/

type Alamat struct {
	Kota, Provinsi, Negara string
}

func changeAlamatWithoutPointer(alamat Alamat) {
	alamat.Negara = "Indonesia"
}

func changeAlamatWitPointer(alamat *Alamat) {
	alamat.Negara = "Indonesia"
}

func main() {
	alamat := Alamat{"Kupang", "NTT", ""}
	changeAlamatWithoutPointer(alamat)
	fmt.Println(alamat) // tidak berubah sama sekali

	alamat2 := Alamat{"Sabu", "NTT", ""}
	changeAlamatWitPointer(&alamat2)
	fmt.Println(alamat2) //berubah

	alamat3 := new(Alamat)
	changeAlamatWitPointer(alamat3)
	fmt.Println(alamat3) //berubah
}
