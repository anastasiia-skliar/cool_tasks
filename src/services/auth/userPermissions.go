package auth

import (
	"github.com/Nastya-Kruglikova/cool_tasks/src/database"
	"github.com/Nastya-Kruglikova/cool_tasks/src/models"
	"github.com/Nastya-Kruglikova/cool_tasks/src/services/common"
	"net/http"
)

var CheckPermission = func(userSession string, requiredRole string, itemOwner string) bool {
	switch requiredRole {
	case "owner":
		return isOwner(userSession, itemOwner)
	case "admin":
		return isAdmin(userSession)
	}
	return false
}

var IsAdmin = func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	if r.URL.Path == "/v1/users" {
		session, err := GetSessionIDFromRequest(w, r)
		if err != nil {
			return
		}
		if isAdmin(session) == false {
			common.SendError(w, r, 400, "ERROR: Not admin", err)
			return
		}
		next(w, r)
		return
	}
	next(w, r)
}

var isOwner = func(session string, itemOwner string) bool {
	sessionLogin, err := database.Cache.Get(session).Result()
	if err != nil {
		return false
	}
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
	sessionLogin, err := database.Cache.Get(session).Result()
	if err != nil {
		return false
	}
	if user, err := models.GetUserByLogin(sessionLogin); err == nil && user.Role == "Admin" {
		return true
	}
	return false
}
