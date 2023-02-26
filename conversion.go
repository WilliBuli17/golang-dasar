package main

import "fmt"

func main() {
	/*
	* Di Go-Lang kadang kita butuh melakukan konversi tipe data dari satu tipe ke tipe lain
	* Misal kita ingin mengkonversi tipe data int32 ke int63, dan lain-lain
	* Krtika konversi tolong perhatikan batas minimum dan maximum dari tipe data barunya yuang ingin dikonversi
	 */

	var nilai32 int32 = 123456
	var nilai64 int64 = int64(nilai32)
	var nilai16 int16 = int16(nilai32)

	fmt.Println(nilai32) // 123456
	fmt.Println(nilai64) // 123456
	fmt.Println(nilai16) // -7616 --> ini karena 123456 melampaui limit maximum dari int16

	var name = "Willi Buli"
	var e byte = name[0]
	var eString = string(e)

	fmt.Println(name)
	fmt.Println(e)
	fmt.Println(eString)
}
