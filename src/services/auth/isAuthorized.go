package auth

import (
	"github.com/Nastya-Kruglikova/cool_tasks/src/services/common"
	"log"
	"net/http"
	"github.com/Nastya-Kruglikova/cool_tasks/src/database"
)

//Start Mocked func that check is the key exist on redis and return true is exist
var IsExistRedis = func(key string) bool {

	var s = database.Cache
	defer s.Close()

	redisKey := "6c3a65d23c5f26fc529f6c5ce01a6b31"

	s.Set(redisKey, "",0)

	if s.Exists(key).String() != "" {
		return true
	}
	return false
}

//End Mocked func

func IsAuthorized(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	if r.URL.Path == "/v1/login" {
		next(w, r)
		return
	}
	userSession, err := r.Cookie("user_session") //get value from user_session key from cookie

	if err != nil {
		log.Println(err, "ERROR: Can't get cookies")
		common.SendError(w, r, 400, "ERROR: Can't get cookies", err)
		return
	}

	if IsExistRedis(userSession.Value) {
		next(w, r)
	} else {
		log.Println(err, "ERROR: Not Authorized")
		common.SendError(w, r, 401, "ERROR: Not Authorized", err)
	}

}
