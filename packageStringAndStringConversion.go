package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*
Package strings adalah package yang berisikan function-function untuk memanipulasi tipe data String
Ada banyak sekali function yang bisa kita gunakan
https://golang.org/pkg/strings/

	Function									Kegunaan
	strings.Trim(string, cutset)				Memotong cutset di awal dan akhir string
	strings.ToLower(string)						Membuat semua karakter string menjadi lower case
	strings.ToUpper(string)						Membuat semua karakter string menjadi upper case
	strings.Split(string, separator)			Memotong string berdasarkan separator
	strings.Contains(string, search)			Mengecek apakah string mengandung string lain
	strings.ReplaceAll(string, from, to)		Mengubah semua string dari from ke to
*/

func main() {
	fmt.Println(strings.Contains("Willi Buli", "Willi"))
	fmt.Println(strings.Split("Willi Buli", " "))
	fmt.Println(strings.ToLower("Willi Buli"))
	fmt.Println(strings.ToUpper("Willi Buli"))
	fmt.Println(strings.Trim("              Willi Buli             ", " "))
	fmt.Println(strings.ReplaceAll("Willi Willi Willi Ama Buli Willi Ama", "Willi", "AWB"))

	/*
		Sebelumnya kita sudah belajar cara konversi tipe data, misal dari int32 ke int34
		Bagaimana jika kita butuh melakukan konversi yang tipe datanya berbeda? Misal dari int ke string, atau sebaliknya
		Hal tersebut bisa kita lakukan dengan bantuan package strconv (string conversion)
		https://golang.org/pkg/strconv/

		Function 							Kegunaan
		strconv.parseBool(string)			Mengubah string ke bool
		strconv.parseFloat(string)			Mengubah string ke float
		strconv.parseInt(string)			Mengubah string ke int64
		strconv.FormatBool(bool)			Mengubah bool ke string
		strconv.FormatFloat(float, … )		Mengubah float64 ke string
		strconv.FormatInt(int, … )			Mengubah int64 ke string
	*/

	boolean, err := strconv.ParseBool("true")

	if err == nil {
		fmt.Println(boolean)
	} else {
		fmt.Println("Error", err.Error())
	}

	number, err := strconv.ParseInt("1000000", 10, 64)
	if err == nil {
		fmt.Println(number)
	} else {
		fmt.Println("Error", err.Error())
	}

	decimal := strconv.FormatInt(1000000, 10)
	biner := strconv.FormatInt(1000000, 2)
	ocrtal := strconv.FormatInt(1000000, 8)
	hexa := strconv.FormatInt(1000000, 16)
	fmt.Println(decimal)
	fmt.Println(biner)
	fmt.Println(ocrtal)
	fmt.Println(hexa)

	valueInt, _ := strconv.Atoi("50000000")
	valueString := strconv.Itoa(50000000)
	fmt.Println(valueInt)
	fmt.Println(valueString)
}
