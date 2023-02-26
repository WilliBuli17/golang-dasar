package main

import "fmt"

func main() {
	/*
			Tipe data Slice adalah potongan dari data Array
			Slice mirip dengan Array, yang membedakan adalah ukuran Slice bisa berubah
			Slide dan Array selalu terkoneksi, dimana Slice adalah data yang mengakses sebagian atau seluruh data di Array

			Tipe Data Slice memiliki 3 data, yaitu pointer, length dan capacity
			Pointer adalah penunjuk data pertama di array para slice
			Length adalah panjang dari slice, dan
			Capacity adalah kapasitas dari slice, dimana length tidak boleh lebih dari capacity

		 	Membuat Slice 	|| 	Keterangan
			array[low:high]		Membuat slice dari array dimulai index low sampai index sebelum high
			array[low:]			Membuat slide dari array dimulai index low sampai index akhir di array
			array[:high]		Membuat slice dari array dimulai index 0 sampai index sebelum high
			array[:]			Membuat slice dari array dimulai index 0 sampai index akhir di array

			Operasi									||	Keterangan
			len(slice)									Untuk mendapatkan panjang
			cap(slice)									Untuk mendapat kapasitas
			append(slice, data)							Membuat slice baru dengan menambah data ke posisi terakhir slice, jika kapasitas sudah penuh, maka akan membuat array baru
			make([]TypeData, length, capacity)			Membuat slice baru
			copy(destination, source)					Menyalin slice dari source ke destination

			Saat membuat Array, kita harus berhati-hati, jika salah, maka yang kita buat bukanlah Array, melainkan Slice
	*/

	var months = [...]string{
		"Jan",
		"Feb",
		"Mar",
		"Apr",
		"Mei",
		"Jun",
		"Jul",
		"Ags",
		"Sep",
		"Okt",
		"Nov",
		"Des",
	}

	slice1 := months[4:8]
	fmt.Println(slice1)

	slice2 := months[3:]
	fmt.Println(slice2)

	slice3 := months[:9]
	fmt.Println(slice3)

	slice4 := months[:]
	fmt.Println(slice4)

	//months[5] = "Diubah"
	//fmt.Println(slice1)

	//slice1[0] = "Mei Lagi"
	//fmt.Println(months)

	//var slice5 = months[10:]
	var slice5 = months[2:4]
	fmt.Println(slice5)

	var slice6 = append(slice5, "Willi")
	fmt.Println(slice6)
	slice6[1] = "Bukan Desember"
	fmt.Println(slice6)

	fmt.Println(slice5)
	fmt.Println(months)

	newSlice := make([]string, 2, 5)

	newSlice[0] = "Willi"
	newSlice[1] = "Buli"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(copySlice)

	iniArray := [5]int{1, 2, 3, 4, 5}
	iniSlice := []int{1, 2, 3, 4, 5}

	fmt.Println(iniArray)
	fmt.Println(iniSlice)
}
