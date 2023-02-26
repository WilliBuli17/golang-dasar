package main

import "fmt"

func main() {
	/*
		Terdiri dari 2 jenis tipe data
			* Integer
			* Floating Point
	*/

	/*
		Integer
		 Tipe Data 	|| Nilai Minimum 		|| Nilai Maksimum
			int8		-128					127
			int16		-32768					32767
			int32		-2147483648				2147483647
			int64		-9223372036854775808	9223372036854775807
			uint8		0						255
			uint16		0						65535
			uint32		0						4294967295
			uint64		0						18446744073709551615
	*/

	/*
		Floating Point
		 Tipe Data 	|| Nilai Minimum 		|| Nilai Maksimum
			float32		1.18×10^(−38)			3.4×10^(38)
			float64		2.23×10^(−308)			1.80×10^(38)
			complex64	complex numbers with float32 real and imaginary parts.
			complex128	complex numbers with float64 real and imaginary parts.
	*/

	/*
		Alias
		 Tipe Data 	|| Alias untuk
			byte		uint8
			rune		int32
			int			Minimal int32
			uint		Minimal uint32
	*/

	fmt.Println("Satu = ", 1)
	fmt.Println("Dua = ", 2)
	fmt.Println("Tiga Koma Lima = ", 3.5)
}
