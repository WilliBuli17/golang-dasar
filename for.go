package main

import "fmt"

func main() {
	/*
		Dalam bahasa pemrograman, biasanya ada fitur yang bernama perulangan
		Saat ini di go hanya ada satu fitur perulangan adalah for loops
	*/

	counter1 := 1
	for counter1 <= 10 {
		fmt.Println("Perulangan ke", counter1)
		counter1++
	}

	for counter2 := 1; counter2 <= 10; counter2++ {
		fmt.Println("Perulangan ke", counter2)
	}

	slice := []string{"Styephen", "William", "Buli", "Budi", "Joko"}

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	for i, value := range slice {
		fmt.Println("Index", i, "=", value)
	}

	// _ menandakan bahwa variabel itu tidak dibutuhkan
	for _, value := range slice {
		fmt.Println(value)
	}

	person := make(map[string]string)
	person["name"] = "Willi"
	person["title"] = "Gabut"

	// kalau map wajib key
	for key, value := range person {
		fmt.Println(key, "=", value)
	}
}
