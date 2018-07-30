package auth

import (
	"github.com/Nastya-Kruglikova/cool_tasks/src/database"
	"github.com/Nastya-Kruglikova/cool_tasks/src/services/common"
	"log"
	"net/http"
)

//IsExistRedis checks if redis exists
var IsExistRedis = func(key string) bool {
	_, err := database.Cache.Get(key).Result()

	return err == nil
}

//IsAuthorized checks authorization
func IsAuthorized(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	if r.URL.Path == "/v1/login" {
		next(w, r)
		return
	}

	userSession, err := GetSessionIDFromRequest(w,r)
	if err != nil {
		return
	}

	if IsExistRedis(userSession) {
		next(w, r)
	} else {
		log.Println(err, "ERROR: Not Authorized")
		common.SendError(w, r, 401, "ERROR: Not Authorized", err)
	}
}
