package auth

import (
	"github.com/Nastya-Kruglikova/cool_tasks/src/services/common"
	"log"
	"net/http"
)

var GetSessionIDFromRequest= func(w http.ResponseWriter, r *http.Request) (string, error) {
	session, err := r.Cookie("user_session")
	if err != nil {
		log.Println(err, "ERROR: Can't get cookies")
		common.SendError(w, r, 400, "ERROR: Can't get cookies", err)
		return "", err
	}
	return session.Value, err
}

func MockedGetSession(sessionID string, err error) {
	GetSessionIDFromRequest = func(w http.ResponseWriter, r *http.Request) (string, error) {
		return sessionID, err
	}
}
