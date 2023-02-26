package main

import "fmt"

/*
Secara default di Go-Lang semua variable itu di passing by value, bukan by reference
Artinya, jika kita mengirim sebuah variable ke dalam function, method atau variable lain, sebenarnya yang dikirim adalah duplikasi value nya

Pointer adalah kemampuan membuat reference ke lokasi data di memory yang sama, tanpa menduplikasi data yang sudah ada
Sederhananya, dengan kemampuan pointer, kita bisa membuat pass by reference (kebaikan dari passing by value)
*/

type Address struct {
	City, Province, Country string
}

func main() {
	address1 := Address{
		City:     "Subang",
		Province: "Jawa Barat",
		Country:  "Indonesia",
	}

	address2 := address1

	address2.City = "Bandung"

	fmt.Println(address1) // address1 tidak berubah sama sekali
	fmt.Println(address2)

	/*
		Untuk membuat sebuah variable dengan nilai pointer ke variable yang lain, kita bisa menggunakan operator & diikuti dengan nama variable nya
	*/
	address3 := Address{
		City:     "Subang",
		Province: "Jawa Barat",
		Country:  "Indonesia",
	}

	address4 := &address3

	address4.City = "Bandung"

	fmt.Println(address3)
	fmt.Println(address4) // ada & saat print karena menandakan bahwa itu hasil pointer

	/*
		Saat kita mengubah variable pointer, maka yang berubah hanya variable tersebut.
		Semua variable yang mengacu ke data yang sama tidak akan berubah
		Jika kita ingin mengubah seluruh variable yang mengacu ke data tersebut, kita bisa menggunakan operator *
	*/

	address5 := Address{
		City:     "Subang",
		Province: "Jawa Barat",
		Country:  "Indonesia",
	}

	address6 := &address5

	address6.City = "Bandung"

	//address6 = &Address{
	//	City:     "Jakarta",
	//	Province: "DKI Jakarta",
	//	Country:  "Indonesia",
	//} // YANG I NI CUMA MERUBAH DI ADDRESS6 SAJA, ADDRESS5 TIDAK IKUT BERUBAH

	*address6 = Address{
		City:     "Jakarta",
		Province: "DKI Jakarta",
		Country:  "Indonesia",
	}

	address7 := &address5 // walaupun di deklarasikan akhir, dia tetap akan ikut berubah jika menggunakan pointer

	fmt.Println(address5)
	fmt.Println(address6) // ada & saat print karena menandakan bahwa itu hasil pointer
	fmt.Println(address7) // ada & saat print karena menandakan bahwa itu hasil pointer

	/*
		Sebelumnya untuk membuat pointer dengan menggunakan operator &
		Go-Lang juga memiliki function new yang bisa digunakan untuk membuat pointer
		Namun function new hanya mengembalikan pointer ke data kosong, artinya tidak ada data awal
	*/

	alamat1 := new(Address)
	alamat2 := alamat1

	alamat2.Country = "Indonesia"

	fmt.Println(alamat1) // ini juga akan berubah
	fmt.Println(alamat2)
}
