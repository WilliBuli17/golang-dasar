package main

import (
	"flag"
	"fmt"
)

/*
Package flag berisikan fungsionalitas untuk memparsing command line argument
https://golang.org/pkg/flag/
*/

func main() {
	host := flag.String("host", "localhost", "Put your database host")
	username := flag.String("username", "root", "Put your database username")
	password := flag.String("password", "root", "Put your database password")

	flag.Parse() // wajib di panggil agar bisa di ubah valuenya

	fmt.Println(*host, *username, *password) // gunakan tanda bintang untuk mengambil value -> karena dia pointer
}
