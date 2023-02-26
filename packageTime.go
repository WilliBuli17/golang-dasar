package main

import (
	"fmt"
	"time"
)

/*
Package time adalah package yang berisikan fungsionalitas untuk management waktu di Go-Lang
https://golang.org/pkg/time/

	Function						Kegunaan
	time.Now()						Untuk mendapatkan waktu saat ini
	time.Date(...)					Untuk membuat waktu
	time.Parse(layout, string)		Untuk memparsing waktu dari string
*/

func main() {
	now := time.Now()
	fmt.Println(now.Local())
	fmt.Println(now.Local().Year())
	fmt.Println(now.Local().Month())
	fmt.Println(now.Local().Day())
	fmt.Println(now.Local().Hour())
	fmt.Println(now.Local().Minute())
	fmt.Println(now.Local().Second())
	fmt.Println(now.Local().Nanosecond())

	utc := time.Date(2009, time.February, 26, 10, 10, 10, 10, time.UTC)
	fmt.Println(utc.Local())

	layout := "2006-01-02"
	parse, _ := time.Parse(layout, "1990-03-20")
	fmt.Print(parse)

	layout2 := time.RFC3339
	parse2, _ := time.Parse(layout2, "1990-03-20T15:04:05Z")
	fmt.Print(parse2)
}
