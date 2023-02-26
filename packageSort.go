package main

import (
	"fmt"
	"sort"
)

/*
Package sort adalah package yang berisikan utilitas untuk proses pengurutan
Agar data kita bisa diurutkan, kita harus mengimplementasikan kontrak di interface sort.Interface
https://golang.org/pkg/sort/
*/

type User struct {
	Name string
	Age  int
}

type UserSlice []User

func (userSlice UserSlice) Len() int {
	return len(userSlice)
}

func (userSlice UserSlice) Less(i, j int) bool {
	return userSlice[i].Age < userSlice[j].Age
}

func (userSlice UserSlice) Swap(i, j int) {
	userSlice[i], userSlice[j] = userSlice[j], userSlice[i]
}

func main() {
	users := []User{
		{
			Name: "Willi",
			Age:  22,
		},
		{
			Name: "Theo",
			Age:  24,
		},
		{
			Name: "Tama",
			Age:  23,
		},
		{
			Name: "Hari",
			Age:  21,
		},
	}

	fmt.Println(users)
	sort.Sort(UserSlice(users))
	fmt.Println(users)
}
