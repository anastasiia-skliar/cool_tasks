package auth

import (
	"github.com/Nastya-Kruglikova/cool_tasks/src/database"
	"github.com/Nastya-Kruglikova/cool_tasks/src/models"
	"net/http"
)

var CheckPermission = func(r *http.Request, requiredRole string, itemOwner string) bool {
	userSession, _ := r.Cookie("user_session")

	switch requiredRole {
	case "owner":
		return isOwner(userSession.Value, itemOwner)
	case "admin":
		return isAdmin(userSession.Value)
	}
	return false
}

var isOwner = func(session string, itemOwner string) bool {
	sessionLogin, _ := database.Cache.Get(session).Result()
	//error is not nil always
	if sessionLogin == itemOwner {
		return true
	}
	if isAdmin(session) {
		return true
	}
	return false
}

var isAdmin = func(session string) bool {
	sessionLogin, _ := database.Cache.Get(session).Result()
	if user, err := models.GetUserByLogin(sessionLogin); err == nil && user.Role == "Admin" {
		return true
	}
	return false
}
