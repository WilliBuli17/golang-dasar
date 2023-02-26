package main

import "fmt"

func getHello(name string) string {
	if name == "" {
		return "Hallo Bro"
	}

	return "Hello " + name
}

func main() {
	result := getHello("Willi")
	fmt.Println(result)

	result = getHello("")
	fmt.Println(result)
}
