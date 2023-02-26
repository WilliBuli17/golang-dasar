package main

import "fmt"

func main() {
	var name1 = "Willi"
	var name2 = "Buli"

	var value1 = 100
	var value2 = 200

	var result = name1 == name2

	println(result)
	fmt.Println(value1 > value2)
	fmt.Println(value1 < value2)
	fmt.Println(value1 == value2)
	fmt.Println(value1 != value2)

}
