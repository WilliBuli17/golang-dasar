package main

import "fmt"

func main() {
	name := "Styephen"

	switch name {
	case "Willi":
		fmt.Println("Hello", name)
	case "Buli":
		fmt.Println("Hello", name)
	default:
		fmt.Println("Kenalan Donk")
	}

	// proses ini cuma bisa true sama false -> mending if-else      kalau ada kasus gini
	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Nama terlalu panjang")
	case false:
		fmt.Println("Nama sudah benar")
	}

	//ini yang bener -> lebih bagus
	length := len(name)
	switch {
	case length > 10:
		fmt.Println("Nama terlalu panjang")
	case length > 5:
		fmt.Println("Nama lumayan panjang")
	default:
		fmt.Println("Nama sudah benar")
	}
}
