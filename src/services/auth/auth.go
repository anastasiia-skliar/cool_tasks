package auth

import (
	"net/http"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"database/sql"
	"github.com/satori/go.uuid"
	"github.com/alicebob/miniredis"
	"github.com/Nastya-Kruglikova/cool_tasks/src/services/common"
	"errors"
)

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
	newLogin.login = r.Form.Get("Login")
	newLogin.pass = r.Form.Get("Pass")

	newLogin.sessionID = tryLogin(w, r, newLogin)
	if newLogin.sessionID != "" {
		redis.Push(newLogin.sessionID, newLogin.login)

		common.RenderJSON(w, r, newLogin.sessionID)

	}
	common.SendError(w, r, 418, "ERROR: ", errors.New("no result"))

}

func tryLogin(w http.ResponseWriter, r *http.Request, loginUser login) string{
	userPass := db.QueryRow("SELECT ID, password FROM users WHERE login = $1", loginUser.login)
	var usersInDB User
	userPass.Scan(&usersInDB.ID, &usersInDB.Password)
	if loginUser.pass == usersInDB.Password {
		sessionUUID, err:= uuid.FromString(usersInDB.Login)
		if err!=nil {

			common.SendError(w, r, 418, "ERROR: ", err)
			return ""
		}
		loginUser.sessionID = sessionUUID.String()
		return loginUser.sessionID
	}
	return ""
}

func Logout(w http.ResponseWriter, r *http.Request) {
	var newLogout login
	r.ParseForm()
	newLogout.sessionID = r.Form.Get("sessionID")
	redis.Del(newLogout.sessionID)
	common.RenderJSON(w, r, "")
}
