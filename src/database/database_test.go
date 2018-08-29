package database_test

import (
	"github.com/Nastya-Kruglikova/cool_tasks/src/config"
	"github.com/Nastya-Kruglikova/cool_tasks/src/database"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"regexp"
	"testing"
)

type databaseTestCase struct {
	name        string
	isConnected bool
	want        bool
}

func TestSetPostgresConnected(t *testing.T) {
	database.IsPostgresConnected = false
	database.SetPostgresConnected()
	if database.IsPostgresConnected != true {
		t.Error("Method SetPostgresConnected is not switching value to true")
	}
}
func TestSetRedisConnected(t *testing.T) {
	database.IsRedisConnected = false
	database.SetRedisConnected()
	if database.IsRedisConnected != true {
		t.Error("Method SetRedisConnected is not switching value to true")
	}
}
func TestSetupPostgres(t *testing.T) {
	tests := []databaseTestCase{
		{
			name:        "Connect postgres first time",
			isConnected: false,
			want:        false,
		},
		{
			name:        "Connect postgres second time",
			isConnected: true,
			want:        true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			database.IsPostgresConnected = tc.isConnected
			switch tc.want {
			case false:
				//Error skipped because we are checking reusing of connection, not if db is up
				db, _, _ := sqlmock.NewWithDSN(database.DSN(config.Config.Database.PostgreSQL))
				database.DB, _ = database.SetupPostgres(config.Config.Database)
				if tc.want != (database.DB == db) {
					t.Error("Error in test: ", tc.name)
				}
			case true:
				database.DB, _ = database.SetupPostgres(config.Config.Database)
				db, err := database.SetupPostgres(config.Config.Database)
				if err != nil {
					t.Log("Postgres is not runing")
				}
				if database.DB != db {
					t.Error("SetupPostgres is not reusing connection, test failed: ", tc.name)
				}
			}

		})
	}
}

func TestDSNRedis(t *testing.T) {
	urls := []database.RedisInfo{
		database.RedisInfo{
			"127.0.0.1",
			1234,
			"pass",
		},

		database.RedisInfo{
			"1.1.0.3",
			6379,
			"pass",
		},
	}

	re := regexp.MustCompile(`(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5]):[0-9]+$`)

	for _, url := range urls {
		if !re.MatchString(database.DSNRedis(url)) {
			t.Error("DSNRedis is returning false data source name", url)
		}
	}
}

func TestSetupRedis(t *testing.T) {
	tests := []databaseTestCase{
		{
			name:        "Connect redis first time",
			isConnected: false,
			want:        false,
		},
		{
			name:        "Connect redis second time",
			isConnected: true,
			want:        true,
		},
	}

	url := database.RedisInfo{
		"127.0.0.1",
		1234,
		"pass",
	}
	mockUrl := database.RedisInfo{
		"0.0.0.0",
		0,
		"",
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			database.IsRedisConnected = tc.isConnected
			switch tc.want {
			case false:
				//Error skipped because we are checking reusing of connection, not if redis db is up
				db, _ := database.SetupRedis(database.Info{Redis: mockUrl})
				database.Cache, _ = database.SetupRedis(database.Info{Redis: url})
				if tc.want != (database.Cache == db) {
					t.Error("Error in test: ", tc.name)
				}
			case true:
				database.Cache, _ = database.SetupRedis(database.Info{Redis: url})
				db, _ := database.SetupRedis(database.Info{Redis: url})
				if database.Cache != db {
					t.Error("SetupRedis is not reusing connection, test failed: ", tc.name)
				}
			}

		})
	}
}
