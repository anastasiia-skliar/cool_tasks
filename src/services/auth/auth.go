//Package auth provides authentication
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

//Login stores info for logging
type login struct {
	login     string
	pass      string
	sessionID string
}

var Login = func(w http.ResponseWriter, r *http.Request) {
	redis := database.Cache
	var newLogin login
	err := r.ParseForm()
	if err != nil {
		common.SendError(w, r, 400, "ERROR: BadRequest", err)
		log.Println(err)
		return
	}
	newLogin.login = r.Form.Get("login")
	newLogin.pass = r.Form.Get("password")
	userInDB, err := models.GetUserForLogin(newLogin.login, newLogin.pass)
	if err != nil {
		common.SendError(w, r, 401, "ERROR: "+err.Error(), err)
		return
	}
	sessionUUID, err := uuid.NewV1()
	if err != nil {
		common.SendError(w, r, 401, "ERROR: "+err.Error(), err)
		return
	}
	newLogin.sessionID = sessionUUID.String()
	if newLogin.sessionID != "" {
		err := redis.Set(newLogin.sessionID, newLogin.login, time.Hour*4).Err()
		if err != nil {
			log.Println(err)
		}
		newCookie := http.Cookie{Name: "user_session", Value: newLogin.sessionID, Expires: time.Now().Add(time.Hour * 4)}
		http.SetCookie(w, &newCookie)
		common.RenderJSON(w, r, userInDB.ID)
		return
	}
	common.SendError(w, r, 401, "ERROR: Authorization failed", errors.New("Fail to autorize"))
}

//Logout logout User
var Logout = func(w http.ResponseWriter, r *http.Request) {
	userSession, err := r.Cookie("user_session")
	if err != nil {
		log.Println(err)
	}
	database.Cache.Del(userSession.Value)
	common.RenderJSON(w, r, "Success logout")
}
