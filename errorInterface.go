package main

import (
	"errors"
	"fmt"
)

/*
Go-Lang memiliki interface yang digunakan sebagai kontrak untuk membuat error, nama interface nya adalah error

Untuk membuat error, kita tidak perlu manual.
Go-Lang sudah menyediakan library untuk membuat helper secara mudah, yang terdapat di package errors (Package akan kita bahas secara detail di materi tersendiri)
*/

func Pembagian(nilai int, penbagi int) (int, error) {
	if penbagi == 0 {
		return 0, errors.New("Pembagi dengan NOL")
	} else {
		return nilai / penbagi, nil
	}
}

func main() {
	hasil, err := Pembagian(100, 10)

	if err == nil {
		fmt.Println("Hasil,", hasil)
	} else {
		fmt.Println("Error,", err)
	}
}
