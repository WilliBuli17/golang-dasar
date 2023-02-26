package main

import "fmt"

/*
Closure adalah kemampuan sebuah function berinteraksi dengan data-data disekitarnya dalam scope yang sama
Harap gunakan fitur closure ini dengan bijak saat kita membuat aplikasi
*/

func main() {
	name := "Willi"
	counter := 0

	increment := func() {
		name = "Buli"
		fmt.Println("Incrememt")
		counter++
	}

	increment()
	increment()

	fmt.Println(counter)
	fmt.Println(name)
}
