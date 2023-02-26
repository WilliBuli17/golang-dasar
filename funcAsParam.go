package main

import "fmt"

/*
Function tidak hanya bisa kita simpan di dalam variable sebagai value
Namun juga bisa kita gunakan sebagai parameter untuk function lain

Kadang jika function terlalu panjang, agak ribet untuk menuliskannya di dalam parameter
Type Declarations juga bisa digunakan untuk membuat alias function, sehingga akan mempermudah kita menggunakan function sebagai parameter
*/

type Filter func(string) string

func sayHelloWithFilter(name string, filter Filter) {
	nameFiltered := filter(name)
	fmt.Println("Hello", nameFiltered)
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}
}

func main() {
	sayHelloWithFilter("Willi", spamFilter)
	sayHelloWithFilter("Anjing", spamFilter)
}
