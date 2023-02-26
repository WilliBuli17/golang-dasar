package main

import "fmt"

/*
Struct adalah sebuah template data yang digunakan untuk menggabungkan nol atau lebih tipe data lainnya dalam satu kesatuan
Struct biasanya representasi data dalam program aplikasi yang kita buat
Data di struct disimpan dalam field
Sederhananya struct adalah kumpulan dari field
*/

// Selalau diawali dengan huruf besar di tiap kata
type Customer struct {
	Name, Address string
	Age           int
	Married       bool
}

func main() {
	/*
		Struct adalah template data atau prototype data
		Struct tidak bisa langsung digunakan
		Namun kita bisa membuat data/object dari struct yang telah kita buat
	*/

	var willi Customer

	willi.Name = "Willi Buli"
	willi.Address = "Jakarta"
	willi.Age = 22

	fmt.Println(willi)

	//Sebelumnya kita telah membuat data dari struct, namun sebenarnya ada banyak cara yang bisa kita gunakan untuk membuat data dari struct
	dito := Customer{
		Name:    "Dito",
		Address: "Jakarta",
		Age:     22,
	}
	fmt.Println(dito)

	theo := Customer{"Theo", "Jogja", 22, false} // harus mengisi semua field -> kalau tidak didisi maka error
	fmt.Println(theo)
}
