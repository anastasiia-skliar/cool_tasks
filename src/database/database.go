//Package database provides DB connections
package database

import (
	"database/sql"
	"fmt"
	"github.com/go-redis/redis"
	_ "github.com/lib/pq"
	"log"
)

//DB is used for postgres connection
//Cache is used for redis connection
//IsPostgresConnected shows is postgres connected
//IsRedisConnected shows is redis connected
var (
	DB                  *sql.DB
	Cache               *redis.Client
	IsPostgresConnected bool
	IsRedisConnected    bool
)

//Type is the type of database from a Type* constant
type Type string

//Info contains information of connections
type Info struct {
	// Database type
	Type []Type
	// Postgres info if used
	PostgreSQL PostgreSQLInfo
	//Redis info
	Redis RedisInfo
}

//PostgreSQLInfo is the details for the database connection
type PostgreSQLInfo struct {
	Hostname     string
	Port         int
	DatabaseName string
	Username     string
	Password     string
}

//RedisInfo shows redis information
type RedisInfo struct {
	URL      string
	Port     int
	Password string
}

//DSN returns the Data Source Name
func DSN(ci PostgreSQLInfo) string {
	return fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=require",
		ci.Hostname, ci.Port, ci.Username, ci.Password, ci.DatabaseName)
}

//DSNRedis Name for Redis
func DSNRedis(ci RedisInfo) string {
	return fmt.Sprintf("%s:%d", ci.URL, ci.Port)
}

//SetupPostgres used for setup connection to postgres
func SetupPostgres(d Info) (*sql.DB, error) {
	if IsPostgresConnected {
		return DB, nil
	}
	db, err := sql.Open("postgres", DSN(d.PostgreSQL))
	if err != nil {
		log.Println(err)
	}

	err = db.Ping()
	if err != nil {
		log.Println(err)
	}
	SetPostgresConnected()
	return db, err
}

//SetupRedis used for setup connection to redis
func SetupRedis(d Info) (*redis.Client, error) {
	if IsRedisConnected {
		return Cache, nil
	}
	client := redis.NewClient(&redis.Options{
		Addr:     DSNRedis(d.Redis),
		Password: d.Redis.Password, // heroku password set
		DB:       0,                // use default DB
	})
	_, err := client.Ping().Result()
	if err != nil {
		log.Println(err)
	}
	SetRedisConnected()
	return client, err
}

//SetPostgresConnected checks postgres connection
func SetPostgresConnected() {
	IsPostgresConnected = true
}

//SetRedisConnected checks redis connection
func SetRedisConnected() {
	IsRedisConnected = true
}
