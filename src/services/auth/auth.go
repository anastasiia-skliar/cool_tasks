//Package auth provides authentication
package auth

import (
	"encoding/json"
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
	Login     string
	Password  string
	SessionID string
}

var Login = func(w http.ResponseWriter, r *http.Request) {
	redis := database.Cache
	var newLogin login

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newLogin)
	if err != nil {
		common.SendError(w, r, 400, "ERROR: Can't decode JSON POST Body", err)
		log.Println(err)
		return
	}

	userInDB, err := models.GetUserForLogin(newLogin.Login, newLogin.Password)
	if err != nil {
		common.SendError(w, r, 401, "ERROR: "+err.Error(), err)
		return
	}

	sessionUUID, err := uuid.NewV1()
	if err != nil {
		common.SendError(w, r, 401, "ERROR: "+err.Error(), err)
		return
	}

	newLogin.SessionID = sessionUUID.String()
	if newLogin.SessionID != "" {
		err := redis.Set(newLogin.SessionID, newLogin.Login, time.Hour*4).Err()
		if err != nil {
			log.Println(err)
		}
		newCookie := http.Cookie{Name: "user_session", Value: newLogin.SessionID, Expires: time.Now().Add(time.Hour * 4)}
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
