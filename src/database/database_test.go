package database_test

import (
	"github.com/Nastya-Kruglikova/cool_tasks/src/config"
	"github.com/Nastya-Kruglikova/cool_tasks/src/database"
	"github.com/rafaeljusto/redigomock"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"testing"
)

func TestSetupPostgres(t *testing.T) {
	database.SetPostgresConnected()
	db, _, _ := sqlmock.NewWithDSN(database.DSN(config.Config.Database.PostgreSQL))
	database.DB = db
	if database.IsPostgresConnected == true {
		db2, err := database.SetupPostgres(config.Config.Database)
		if db2 != database.DB {
			t.Fatal(err)
		}
	}

}

func TestSetupRedis(t *testing.T) {
	database.SetRedisConnected()
	conn := redigomock.NewConn()
	database.Cache = conn
	if database.IsRedisConnected == true {
		conn, err := database.SetupRedis(config.Config.Database)
		if conn != database.Cache {
			t.Fatal(err)
		}
	}
}
