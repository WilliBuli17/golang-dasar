package main

import "fmt"

func main() {
	/*
	* String ada tipe data kumpulan karakter
	* Jumlah karakter di dalam String bisa nol sampai tidak terhingga
	* Tipe data String di Go-Lang direpresentasikan dengan kata kunci string
	* Nilai data String di Go-Lang selalu diawali dengan karakter “ (petik dua) dan diakhiri dengan karakter “ (petik dua)
	 */

	fmt.Println("Willi")
	fmt.Println("Willi Buli")
	fmt.Println("Styephen William Buli")

	/*
		Function	 	|| Keterangan
		len(“string”)		Menghitung jumlah karakter di String
		“string”[number]	Mengambil karakter pada posisi yang ditentukan --> Balikannya byte, misalnya 87 adalah bytenya "W" dan 116 adalah bytenya "t"
	*/

	fmt.Println(len("Willi"))
	fmt.Println("Willi Buli"[0])
	fmt.Println("Styephen William Buli"[1])
}
