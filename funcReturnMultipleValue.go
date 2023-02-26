package main

func getFullName() (string, string, string) {
	return "Styephen", "William", "Buli"
}

func main() {
	value1, value2, value3 := getFullName()

	firstName := value1
	lastName := value2 + " " + value3

	println(firstName, lastName)

	value4, value5, _ := getFullName()

	println(value4, value5)
}
