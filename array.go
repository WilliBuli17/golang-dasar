package main

import "fmt"

func main() {

	var names [3]string

	names[0] = "Styephen"
	names[1] = "William"
	names[2] = "Buli"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	var values = [3]int{
		80,
		85,
		90,
	}

	fmt.Println(values)
	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])

	values[0] = 100
	fmt.Println(values)

	fmt.Println(len(names))
	fmt.Println(len(values))

	var lagi [10]string
	fmt.Println(len(lagi))
}
