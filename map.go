package main

import "fmt"

func main() {
	/*
		Pada Array atau Slice, untuk mengakses data, kita menggunakan index Number dimulai dari 0
		Map adalah tipe data lain yang berisikan kumpulan data yang sama, namun kita bisa menentukan jenis tipe data index yang akan kita gunakan
		Sederhananya, Map adalah tipe data kumpulan key-value (kata kunci - nilai), dimana kata kuncinya bersifat unik, tidak boleh sama
		Berbeda dengan Array dan Slice, jumlah data yang kita masukkan ke dalam Map boleh sebanyak-banyaknya, asalkan kata kunci nya berbeda, jika kita gunakan kata kunci sama, maka secara otomatis data sebelumnya akan diganti dengan data baru

		Operasi							||	Keterangan
		len(map)							Untuk mendapatkan jumlah data di map
		map[key]							Mengambil data di map dengan key
		map[key] = value					Mengubah data di map dengan key
		make(map[TypeKey]TypeValue)			Membuat map baru
		delete(map, key)					Menghapus data di map dengan key
	*/
	person := map[string]string{
		"name":    "Willi",
		"address": "Hutan",
	}

	person["title"] = "Gabut"

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(person["title"])

	var book = make(map[string]string)
	book["title"] = "Belajar Golang"
	book["author"] = "Willi"
	book["ups"] = "Salah"
	fmt.Println(book)
	fmt.Println(len(book))

	delete(book, "ups")

	fmt.Println(book)
	fmt.Println(len(book))
}
