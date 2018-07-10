package auth

import (
	"errors"
	"github.com/Nastya-Kruglikova/cool_tasks/src/models"
	"github.com/Nastya-Kruglikova/cool_tasks/src/services/common"
	"github.com/satori/go.uuid"
	"net/http"
	"time"
	"fmt"
	"github.com/Nastya-Kruglikova/cool_tasks/src/database"
	"log"
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

func Login(w http.ResponseWriter, r *http.Request) {
	 GetUserByLogin := models.GetUserByLogin
	redis := database.Cache
	/*userSession, err := r.Cookie("user_session")
	if err == nil {
		if redis.Get(userSession.Value) != nil {
			login := redis.Get(userSession.Value)
			redis.Del(userSession.Value)
			redis.Set(userSession.Value, login, time.Hour*4)

			newCookie := http.Cookie{Name: "user_session", Value: userSession.Value, Expires: time.Now().Add(time.Hour * 4)}
			http.SetCookie(w, &newCookie)

			common.RenderJSON(w, r, userSession.Value)
			return
		}
	}*/
	var newLogin login

	r.ParseForm()
	newLogin.login = r.Form.Get("login")
	newLogin.pass = r.Form.Get("password")

	var userInDB models.User
	fmt.Println(newLogin)
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

		fmt.Println(newLogin)
		err:=redis.Set(newLogin.sessionID, newLogin.login, time.Hour*4).Err()
		if err!= nil {
			log.Println(err)
		}
		//fmt.Println(err)
		fmt.Println("%%%%%%%%%")
		val,err := database.Cache.Get(newLogin.sessionID).Result()
		fmt.Println("Value:=" + val)

		newCookie := http.Cookie{Name: "user_session", Value: newLogin.sessionID, Expires: time.Now().Add(time.Hour * 4)}
		http.SetCookie(w, &newCookie)

		common.RenderJSON(w, r, newLogin.sessionID)
		return
	}
	common.SendError(w, r, 401, "ERROR: ", errors.New("Fail to autorize"))

}

func Logout(w http.ResponseWriter, r *http.Request) {
	userSession, _ := r.Cookie("user_session")
	database.Cache.Del(userSession.Value)
	common.RenderJSON(w, r, "Success logout")

}
