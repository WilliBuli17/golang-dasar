package helper

/*
Package adalah tempat yang bisa digunakan untuk mengorganisir kode program yang kita buat di Go-Lang
Dengan menggunakan package, kita bisa merapikan kode program yang kita buat
Package sendiri sebenarnya hanya direktori folder di sistem operasi kita
*/

/*
Di bahasa pemrograman lain, biasanya ada kata kunci yang bisa digunakan untuk menentukan access modifier terhadap suatu function atau variable
Di Go-Lang, untuk menentukan access modifier, cukup dengan nama function atau variable
Jika nama nya diawali dengan hurup besar, maka artinya bisa diakses dari package lain, jika dimulai dengan hurup kecil, artinya tidak bisa diakses dari package lain

methodnya harus diawali dengan huruf besar jika ingin bisa diakses dari luar
*/

var version = "1.0.0"
var Aplication = "golang"

func SayHai(name string) string {
	return "Hallo " + name
}

func sayGoodbye(name string) string {
	return "Bye " + name
}
