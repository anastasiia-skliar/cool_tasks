package auth

import (
	"errors"
	"github.com/Nastya-Kruglikova/cool_tasks/src/database"
	"github.com/Nastya-Kruglikova/cool_tasks/src/models"
	"github.com/Nastya-Kruglikova/cool_tasks/src/services/common"
	"github.com/satori/go.uuid"
	"log"
	"net/http"
	"time"
)

type login struct {
	login     string
	pass      string
	sessionID string
}

var Login = func(w http.ResponseWriter, r *http.Request) {
	GetUserByLogin := models.GetUserByLogin
	redis := database.Cache
	var newLogin login

	r.ParseForm()
	newLogin.login = r.Form.Get("login")
	newLogin.pass = r.Form.Get("password")

	userInDB, er := GetUserByLogin(newLogin.login)
	if er != nil {
		common.SendError(w, r, 401, "ERROR: ", er)
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

		err := redis.Set(newLogin.sessionID, newLogin.login, time.Hour*4).Err()
		if err != nil {
			log.Println(err)
		}

		newCookie := http.Cookie{Name: "user_session", Value: newLogin.sessionID, Expires: time.Now().Add(time.Hour * 4)}
		http.SetCookie(w, &newCookie)

		common.RenderJSON(w, r, newLogin.sessionID)
		return
	}
	common.SendError(w, r, 401, "ERROR: ", errors.New("Fail to autorize"))

}

var Logout = func(w http.ResponseWriter, r *http.Request) {
	userSession, _ := r.Cookie("user_session")
	database.Cache.Del(userSession.Value)
	common.RenderJSON(w, r, "Success logout")
}
