package main

import "fmt"

/*
Struct adalah tipe data seperti tipe data lainnya, dia bisa digunakan sebagai parameter untuk function
Namun jika kita ingin menambahkan method ke dalam structs, sehingga seakan-akan sebuah struct memiliki function
Method adalah function
*/
type Customer2 struct {
	Name, Address string
	Age           int
	Married       bool
}

func (customer Customer2) sayHello(name string) {
	fmt.Println("Hallo", name, ", My Name is", customer.Name)
}

func main() {
	willi := Customer2{Name: "Willi"}
	willi.sayHello("Kamu")
}
