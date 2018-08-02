package auth

import (
	"github.com/Nastya-Kruglikova/cool_tasks/src/database"
	"github.com/Nastya-Kruglikova/cool_tasks/src/models"
	"github.com/Nastya-Kruglikova/cool_tasks/src/services/common"
	"net/http"
)

var SpecialPermissions map[string][]string

const AdminRole = "admin"
const Owner = "owner"
const NotOwnerResponse = "User is not object owner"

func init() {
	SpecialPermissions = make(map[string][]string)
	SpecialPermissions["/v1/users"] = []string{AdminRole}
}

var CheckPermission = func(userSession string, requiredRole string, itemOwner string) bool {
	switch requiredRole {
	case Owner:
		return isOwner(userSession, itemOwner) || isAdmin(userSession)
	case AdminRole:
		return isAdmin(userSession)
	}
	return false
}

var AccessPermission = func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	rolesRequired := SpecialPermissions[r.URL.Path]
	if rolesRequired == nil {
		next(w, r)
		return
	}
	for _, v := range rolesRequired {
		switch v {
		case AdminRole:
			session, err := GetSessionIDFromRequest(w, r)
			if err != nil {
				return
			}
			if isAdmin(session) == false {
				common.SendError(w, r, 403, "ERROR: Not admin", err)
				return
			}
			next(w, r)
			return
		default:
			next(w, r)
			return
		}
	}
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
	return false
}

var isAdmin = func(session string) bool {
	sessionLogin, err := database.Cache.Get(session).Result()
	if err != nil {
		return false
	}
	if user, err := models.GetUserByLogin(sessionLogin); err == nil && user.Role == AdminRole {
		return true
	}
	return false
}
