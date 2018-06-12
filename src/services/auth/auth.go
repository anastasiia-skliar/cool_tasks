package auth

import (
	"net/http"
	"encoding/json"
)

type login struct {
	id string
	login string
	pass string
	sessionID string
}
func Login(w http.ResponseWriter, r *http.Request) {

	var newLogin login

	r.ParseForm()
	newLogin.login = r.Form.Get("Login")
	newLogin.pass=r.Form.Get("Pass")

newLogin.sessionID=tryLogin(newLogin)
if newLogin.sessionID != ""{
	w.Header().Set("Content-Type", "application/json")
	jsonResp, _ :=json.Marshal(newLogin.sessionID)
	w.Write(jsonResp)
}

}

func tryLogin (loginUser login) string{
	usersInDB, err:=db.Query("SELECT * FROM users WHERE login = $1",loginUser.login)
	if err!=nil{
		return ""
	}
	if  loginUser.pass ==usersInDB.pass{
		redis.CreateSession(userInDB.id)
		return userInDB.id
	}
	return ""
}

func Logout(w http.ResponseWriter, r *http.Request)  {
	var newLogout login

	r.ParseForm()
	newLogout.sessionID = r.Form.Get("sessionID")
	redis.FinishSession(newLogout.sessionID)
}