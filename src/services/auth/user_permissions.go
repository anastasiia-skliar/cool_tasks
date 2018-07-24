package auth

import (
	"github.com/Nastya-Kruglikova/cool_tasks/src/database"
	"github.com/Nastya-Kruglikova/cool_tasks/src/models"
)

var IsOwner = func (session string, item_owner string) bool{
	sessionLogin, _:=database.Cache.Get(session).Result()
	//error is not nill always
	if sessionLogin==item_owner{
		return true
	}
	if IsAdmin(session){

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
