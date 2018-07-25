package auth

import (
	"github.com/Nastya-Kruglikova/cool_tasks/src/database"
	"github.com/Nastya-Kruglikova/cool_tasks/src/models"
	"net/http"
)

var CheckPermission = func(r *http.Request, requiredRole string, itemOwner string) bool{
	userSession, _ := r.Cookie("user_session")

	switch requiredRole {
	case "owner":
		return IsOwner(userSession.Value, itemOwner)
	case "admin":
		return IsAdmin(userSession.Value)
	}
	return false
}

var IsOwner = func (session string, itemOwner string) bool{
	sessionLogin, _:=database.Cache.Get(session).Result()
	//error is not nil always
	if sessionLogin==itemOwner{
		return true
	}
	if IsAdmin(session){
return true
	}
	return false
}

var IsAdmin = func(session string) bool{
	sessionLogin, _:=database.Cache.Get(session).Result()
	if user, err:=models.GetUserByLogin(sessionLogin); err==nil && user.Role=="Admin"{
		return true
	}
	return false
}
