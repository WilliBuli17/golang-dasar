package main

import (
	"fmt"
	"os"
)

/*
Go-Lang telah menyediakan banyak sekali package bawaan, salah satunya adalah package os
Package os berisikan fungsionalitas untuk mengakses fitur sistem operasi secara independen (bisa digunakan  disemua sistem operasi)
https://golang.org/pkg/os/
*/

func main() {
	args := os.Args
	fmt.Println("Argument : ")
	fmt.Println(args)

	hostname, err := os.Hostname()

	if err == nil {
		fmt.Println(hostname)
	} else {
		fmt.Println("Error,", err.Error())
	}
}
