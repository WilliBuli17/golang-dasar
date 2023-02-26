package main

import "fmt"

func sayHelloWithFlter(firstname string, lastname string) {
	fmt.Println("Hello", firstname, lastname)
}

func main() {
	sayHelloWithFlter("Willi", "Buli")
}
