package main

import "fmt"

/*
Sebelumnya setiap membuat function, kita akan selalu memberikan sebuah nama pada function tersebut
Namun kadang ada kalanya lebih mudah membuat function secara langsung di variable atau parameter tanpa harus membuat function terlebih dahulu
Hal tersebut dinamakan anonymous function, atau function tanpa nama
*/

type Blackkist func(string) bool

func registerUser(name string, blacklist Blackkist) {
	if blacklist(name) {
		fmt.Println("You are blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}

//ini jika tidak anonymous
//func blacklistAdmin(name string) bool {
//	return name == "admin"
//}
//
//func blacklistRoot(name string) bool {
//	return name == "root"
//}

func main() {

	//cara define dulu
	blacklist := func(name string) bool {
		return name == "admin"
	}

	registerUser("admin", blacklist)
	registerUser("eko", blacklist)

	//cara langsung
	registerUser("root", func(name string) bool {
		return name == "root"
	})
	registerUser("eko", func(name string) bool {
		return name == "root"
	})
}
