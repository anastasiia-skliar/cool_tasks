package auth

import (
	"errors"
	"github.com/Nastya-Kruglikova/cool_tasks/src/models"
	"github.com/Nastya-Kruglikova/cool_tasks/src/services/common"
	"github.com/alicebob/miniredis"
	"github.com/satori/go.uuid"
	"net/http"
	"time"
)

var GetUserByLogin = models.GetUserByLogin

var redis *miniredis.Miniredis

func init() {
	redis, _ = miniredis.Run()
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

func Login(w http.ResponseWriter, r *http.Request) {

	var newLogin login

	r.ParseForm()
	newLogin.login = r.Form.Get("login")
	newLogin.pass = r.Form.Get("password")

	var userInDB models.User
	userInDB, err := GetUserByLogin(newLogin.login)
	if err != nil {
		common.SendError(w, r, 401, "ERROR: ", err)
		return

	}

	if newLogin.pass == userInDB.Password {
		sessionUUID, err := uuid.NewV1()
		if err != nil {
			common.SendError(w, r, 401, "ERROR: ", err)
			return
		}
		newLogin.sessionID = sessionUUID.String()
	}
	if newLogin.sessionID != "" {

		redis.Push(newLogin.sessionID, newLogin.login)

		newCookie := http.Cookie{Name: "user_session", Value: newLogin.sessionID, Expires: time.Now().Add(time.Hour * 4)}
		http.SetCookie(w, &newCookie)

		common.RenderJSON(w, r, newLogin.sessionID)
		return
	}
	common.SendError(w, r, 401, "ERROR: ", errors.New("Fail to autorize"))

}

func Logout(w http.ResponseWriter, r *http.Request) {
	userSession, _ := r.Cookie("user_session")
	redis.Del(userSession.Value)
	common.RenderJSON(w, r, "Success logout")

}
