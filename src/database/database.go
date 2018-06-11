package database

import (
	"fmt"
	_ "github.com/lib/pq" //PostgreSQL driver
	"database/sql"
	"github.com/garyburd/redigo/redis"
)

var (
	databases Info
)

// Type is the type of database from a Type* constant
type Type string

type Info struct {
	// Database type
	Type[] Type
	// Postgres info if used
	PostgreSQL PostgreSQLInfo
	//Redis info
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
	return fmt.Sprintf("%s:%d",ci.URL,ci.Port)
}

func SetupPostgres(d Info) (*sql.DB, error) {
	db, err := sql.Open("postgres", DSN(d.PostgreSQL))
	//check if is alive
	err=db.Ping()
	return db,err
}

func SetupRedis(d Info)(redis.Conn,error){
	var err,pool = newPool(d)
	connection := pool.Get()
	return connection,err
}

func newPool(d Info) (error,*redis.Pool) {
	c,err:=redis.Dial("tcp", DSN_Redis(d.Redis))
	return err,&redis.Pool{
		MaxIdle:   80,
		MaxActive: 12000, // max number of connections
		Dial: func() (redis.Conn, error) {
			return c, err
		},
	}
}

// ReadConfig returns the database information
func ReadConfig() Info {
	return databases
}
