package database

import (
	"fmt"
	"log"
	_ "github.com/lib/pq" //PostgreSQL driver
	"database/sql"
)

var (
	SQL *sql.DB
	databases Info
)

// Type is the type of database from a Type* constant
type Type string

const (
	TypePostgreSQL Type = "PostgreSQL"
	//TODO: add REDIS support
)

type Info struct {
	// Database type
	Type Type
	// Postgres info if used
	PostgreSQL PostgreSQLInfo
	//TODO: add REDIS support

}

// PostgreSQLInfo is the details for the database connection
type PostgreSQLInfo struct {
	Hostname  string
	Port  int
	DatabaseName string
	Username string
	Password string
}


// DSN returns the Data Source Name
func DSN(ci PostgreSQLInfo) string {
	return fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		ci.Hostname, ci.Port, ci.Username, ci.Password, ci.DatabaseName)
}

// Connect to the database
func Connect(d Info){
	var err error

	// Store the config
	databases = d

	switch d.Type {

	case TypePostgreSQL:
		// Connect to PostgreSQL
		SQL, err = sql.Open("postgres", DSN(d.PostgreSQL));
		if  err != nil {
			log.Println("SQL Driver Error", err)
		}
		// Check if is alive
		if err = SQL.Ping(); err != nil {
			log.Println("Database Error", err)
		}
		log.Println("Connected to PostgreSQL")
	default:
		log.Println("No registered database in config")
	}
}

// ReadConfig returns the database information
func ReadConfig() Info {
	return databases
}
