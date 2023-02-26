package main

import "fmt"

func main() {

	var a = 5
	var b = 10
	var c = a * b
	var d = 10 + 10
	fmt.Println(c)
	fmt.Println(d)

	//Augmented Assignments
	a += c
	fmt.Println(a)

	//Unary Operator
	b++
	d--
	fmt.Println(b)
	fmt.Println(d)
}
