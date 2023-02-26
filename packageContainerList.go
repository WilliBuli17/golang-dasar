package main

import (
	"container/list"
	"fmt"
)

/*
Package container/list adalah implementasi struktur data double linked list di Go-Lang
https://golang.org/pkg/container/list/
*/

func main() {
	data := list.New()
	data.PushBack("Styephen")
	data.PushBack("William")
	data.PushBack("Buli")
	data.PushFront("Ama")

	fmt.Println(data.Back().Value)
	fmt.Println(data.Back().Prev().Value)
	fmt.Println(data.Back().Prev().Prev().Value)
	fmt.Println(data.Back().Prev().Prev().Prev().Value)

	fmt.Println(data.Front().Value)
	fmt.Println(data.Front().Next().Value)
	fmt.Println(data.Front().Next().Next().Value)
	fmt.Println(data.Front().Next().Next().Next().Value)

	for e := data.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

	for e := data.Back(); e != nil; e = e.Prev() {
		fmt.Println(e.Value)
	}
}
