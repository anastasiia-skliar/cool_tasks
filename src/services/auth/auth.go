package auth

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/Nastya-Kruglikova/cool_tasks/src/services/common"
	"github.com/alicebob/miniredis"
	"github.com/satori/go.uuid"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"net/http"
)

func init() {
	db, _, _ = sqlmock.New()
}

type login struct {
	id        uuid.UUID
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
	newLogin.login = r.Form.Get("login")
	newLogin.pass = r.Form.Get("password")

	var usersInDB User
	err := db.QueryRow("SELECT ID, Password FROM Users WHERE Login = $1", newLogin.login).Scan(&(usersInDB.ID), &(usersInDB.Password))
	if err != nil {
		fmt.Println(err)
		return
	}

	if newLogin.pass == usersInDB.Password {
		sessionUUID, err := uuid.NewV1()
		if err != nil {
			fmt.Println("##### ", err)
			common.SendError(w, r, 418, "ERROR: ", err)
			return
		}
		newLogin.sessionID = sessionUUID.String()
	}
	if newLogin.sessionID != "" {
		redis.Push(newLogin.sessionID, newLogin.login)

		common.RenderJSON(w, r, newLogin.sessionID)
		return
	}
	common.SendError(w, r, 418, "ERROR: ", errors.New("no result"))

}

func Logout(w http.ResponseWriter, r *http.Request) {
	var newLogout login
	r.ParseForm()
	newLogout.sessionID = r.Form.Get("sessionID")
	redis.Del(newLogout.sessionID)
	common.RenderJSON(w, r, "")
}
