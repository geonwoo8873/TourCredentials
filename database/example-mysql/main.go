package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type DatabaseConfig struct {
	user     string
	password string
	name     string
	host     string
	port     string
}

func main() {
	var mysql_config = DatabaseConfig{
		user:     os.Getenv("MYSQL_USER"),
		password: os.Getenv("MYSQL_PASSWORD"),
		name:     os.Getenv("MYSQL_DATABASE"),
		host:     os.Getenv("MYSQL_HOST"),
		port:     os.Getenv("MYSQL_PORT"),
	}

	// Open a connection to the MySQL database
	db, err := sql.Open(
		"mysql", mysql_config.user+":"+mysql_config.password+"@tcp("+mysql_config.host+":"+mysql_config.port+")/"+mysql_config.name)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Test the connection
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
}
