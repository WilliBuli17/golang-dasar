package database

var connection string

func init() {
	connection = "PgSQL"
}

func GetDB() string {
	return connection
}
