package database

import (
	"database/sql"
	"fmt"
	"github.com/go-redis/redis"
	_ "github.com/lib/pq"
	"log"
)

var (
	DB                  *sql.DB
	Cache               *redis.Client
	IsPostgresConnected bool
	IsRedisConnected    bool
)

// Type is the type of database from a Type* constant
type Type string

type Info struct {
	// Database type
	Type []Type
	// Postgres info if used
	PostgreSQL PostgreSQLInfo
	//Redis info
	Redis RedisInfo
}

// PostgreSQLInfo is the details for the database connection
type PostgreSQLInfo struct {
	Hostname     string
	Port         int
	DatabaseName string
	Username     string
	Password     string
}
type RedisInfo struct {
	URL  string
	Port int
}

// DSN returns the Data Source Name
func DSN(ci PostgreSQLInfo) string {
	return fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		ci.Hostname, ci.Port, ci.Username, ci.Password, ci.DatabaseName)
}

//Data Source Name for Redis
func DSN_Redis(ci RedisInfo) string {
	return fmt.Sprintf("%s:%d", ci.URL, ci.Port)
}

//Setup Postgres Connection
func SetupPostgres(d Info) (*sql.DB, error) {
	if IsPostgresConnected == true {
		return DB, nil
	}
	db, err := sql.Open("postgres", DSN(d.PostgreSQL))
	//check if is alive
	err = db.Ping()
	if err != nil {
		log.Println(err)
	}
	SetPostgresConnected()
	return db, err
}

func SetupRedis(d Info) (*redis.Client, error) {
	if IsRedisConnected == true {
		return Cache, nil
	}
	client := redis.NewClient(&redis.Options{
		Addr:     DSN_Redis(d.Redis),
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	_, err := client.Ping().Result()
	if err != nil {
		log.Println(err)
	}
	SetRedisConnected()
	return client, err
}

//Sets boolean isPostgresConnected to true
func SetPostgresConnected() {
	IsPostgresConnected = true
}

//Sets boolean isRedisConnected to true
func SetRedisConnected() {
	IsRedisConnected = true
}
