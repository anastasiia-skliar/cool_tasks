package auth

import (
	"encoding/json"
	"net/http"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"database/sql"
	"github.com/satori/go.uuid"
	"github.com/alicebob/miniredis"
)

type login struct {
	id        string
	login     string
	pass      string
	sessionID string
}

type User struct {
	ID       uuid.UUID
	Name     string
	Login    string
	Password string
}
var redis *miniredis.Miniredis
var db *sql.DB
func init() {
    redis, _ = miniredis.Run()
	db, _, _ = sqlmock.New()
}

func Login(w http.ResponseWriter, r *http.Request) {

	var newLogin login

	r.ParseForm()
	newLogin.login = r.Form.Get("Login")
	newLogin.pass = r.Form.Get("Pass")

	newLogin.sessionID = tryLogin(newLogin)
	if newLogin.sessionID != "" {
		w.Header().Set("Content-Type", "application/json")
		jsonResp, _ := json.Marshal(newLogin.sessionID)
		w.Write(jsonResp)
	}

}

func tryLogin(loginUser login) string {
	userPass := db.QueryRow("SELECT ID, password FROM users WHERE login = $1", loginUser.login)
	var usersInDB User
	userPass.Scan(&usersInDB.ID, &usersInDB.Password)
	if loginUser.pass == usersInDB.Password {
		if redis.Exists(  usersInDB.ID.String()){
			redis.Del()
		}
		redis.CreateSession(userInDB.id)
		return userInDB.id
	}
	return ""
}

func Logout(w http.ResponseWriter, r *http.Request) {
	var newLogout login

	r.ParseForm()
	newLogout.sessionID = r.Form.Get("sessionID")
	redis.FinishSession(newLogout.sessionID)
}
