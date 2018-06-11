package auth

import (
	"time"
	"net/http"
)

type login struct {
	id string
	login string
	pass string
}
func Login(w http.ResponseWriter, r *http.Request) {

	var newLogin *login=new(login)

	r.ParseForm()
	newLogin.login = r.Form.Get("Login")
	newLogin.pass=r.Form.Get("Pass")
}

func tryLogin (loginUser login){
	userInDB:=db.search("login",loginUser.login)
	if  userData ==userData.pass{
		redis.CreateSession(userInDB.id)
		return userInDB.id
	}
}

func Logout(logoutUser login)  {
	redis.FinishSession(logoutUser.id)
}