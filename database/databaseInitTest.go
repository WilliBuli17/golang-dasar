package database

import "fmt"

var connect string

func init() {
	fmt.Println("Init dijalankan")
	connect = "MySQL"
}

func GetDatabase() string {
	return connect
}
