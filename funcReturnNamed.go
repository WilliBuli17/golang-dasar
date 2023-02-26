package main

func getCompleteName() (firstName, midleName, lastName string) {
	firstName = "Styephen"
	midleName = "William"
	lastName = "Buli"

	return
}

func main() {
	value1, value2, value3 := getCompleteName()

	println(value1, value2, value3)
}
