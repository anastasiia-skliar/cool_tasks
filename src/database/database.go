package database

import (
	"database/sql"
	"fmt"
	"github.com/garyburd/redigo/redis"
	_ "github.com/lib/pq" //PostgreSQL driver
	"log"
)

var (
	DB                  *sql.DB
	Cache               redis.Conn
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
	if IsPostgresConnected == false {
		db, err := sql.Open("postgres", DSN(d.PostgreSQL))
		//check if is alive
		err = db.Ping()
		SetPostgresConnected()
		log.Println("new connection to postgres")
		return db, err
	}
	var err error
	log.Println("reusing connection postgres connection ")
	return DB, err
}

func SetupRedis(d Info) (redis.Conn, error) {
	if IsRedisConnected == false {
		var err, pool = newPool(d)
		connection := pool.Get()
		SetRedisConnected()
		return connection, err
	}
	var err error
	return Cache, err
}

//New redis connection pool
func newPool(d Info) (error, *redis.Pool) {
	c, err := redis.Dial("tcp", DSN_Redis(d.Redis))
	return err, &redis.Pool{
		MaxIdle:   80,
		MaxActive: 12000, // max number of connections
		Dial: func() (redis.Conn, error) {
			return c, err
		},
	}
}


//Sets boolean isPostgresConnected to true
func SetPostgresConnected() {
	IsPostgresConnected = true
}

//Sets boolean isRedisConnected to true
func SetRedisConnected() {
	IsRedisConnected = true
}
