package main

import "fmt"

func main() {
	var nilaiAkhir = 90
	var absensi = 80

	var lulusNilaiAkhir = nilaiAkhir > 80
	var lulusAbsensi = absensi > 80

	fmt.Println(lulusNilaiAkhir)
	fmt.Println(lulusAbsensi)
	fmt.Println(lulusNilaiAkhir && lulusAbsensi)
	fmt.Println(lulusNilaiAkhir || lulusAbsensi)
	fmt.Println(!lulusNilaiAkhir)
	fmt.Println(!lulusAbsensi)
}
