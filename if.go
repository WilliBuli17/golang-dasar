package main

import "fmt"

func main() {
	name := "Buli"

	if name == "Willi" {
		fmt.Println("Hallo", name)
	} else if name == "Buli" {
		fmt.Println("Hallo", name)
	} else {
		fmt.Println("Gak tau nih")
	}

	var length1 = len(name)
	if length1 > 2 {
		fmt.Println("Ribet 1,", length1)
	}

	if length2 := len(name); length2 > 2 {
		fmt.Println("Ribet 2,", length2)
	}

	if len(name) > 2 {
		fmt.Println("Simple")
	}
}
