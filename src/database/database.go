package database

import (
	"fmt"
	"log"
	_ "github.com/lib/pq" //PostgreSQL driver
	"database/sql"
	"github.com/garyburd/redigo/redis"
)

var (
	RedisPool *redis.Pool // creating redis pool enables us to reuse redigo connections
	SQL *sql.DB
	databases Info
)

// Type is the type of database from a Type* constant
type Type string

const (
	TypePostgreSQL Type = "PostgreSQL"
	TypeRedis Type = "Redis"
)

type Info struct {
	// Database type
	Type[] Type
	// Postgres info if used
	PostgreSQL PostgreSQLInfo
	Redis RedisInfo

}

// PostgreSQLInfo is the details for the database connection
type PostgreSQLInfo struct {
	Hostname  string
	Port  int
	DatabaseName string
	Username string
	Password string
}
type RedisInfo struct {
	URL string
	Port int
}


// DSN returns the Data Source Name
func DSN(ci PostgreSQLInfo) string {
	return fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		ci.Hostname, ci.Port, ci.Username, ci.Password, ci.DatabaseName)
}
func DSN_Redis(ci RedisInfo) string {
	return fmt.Sprintf("%s"+":"+"%d",ci.URL,ci.Port)

}


// Connect to the database
func Connect(d Info){
	var err error

	// Store the config
	databases = d

	for index,_:= range d.Type[:]{
		switch d.Type[index] {
		case TypePostgreSQL:
			// Connect to PostgreSQL
			SQL, err = sql.Open("postgres", DSN(d.PostgreSQL));
			if err != nil {
				log.Println("SQL Driver Error", err)
			}
			// Check if is alive
			if err = SQL.Ping(); err != nil {
				log.Println("Database Error", err)
			} else {
				log.Println("Connected to Postgres")
			}
		case TypeRedis:{
			//redisAddr :=  DSN_Redis(d.Redis)
			RedisPool = &redis.Pool{
				Dial: func() (redis.Conn, error) {
					conn, err := redis.Dial("tcp", DSN_Redis(d.Redis))
					if err==nil{
						log.Println("Connected to Redis")
					}
					return conn, err
				},
			}
		}
		default:
			log.Println("No registered database in config")
		}
	}
}

// ReadConfig returns the database information
func ReadConfig() Info {
	return databases
}
