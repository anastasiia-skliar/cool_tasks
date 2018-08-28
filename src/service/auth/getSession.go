package auth

import (
	"log"
	"net/http"
)

var GetSessionIDFromRequest = func(w http.ResponseWriter, r *http.Request) (string, error) {
	session, err := r.Cookie("user_session")
	if err != nil {
		log.Println(err, "ERROR: Can't get cookies")
		return "", err
	}
	return session.Value, err
}

func MockedGetSession(sessionID string, err error) {
	GetSessionIDFromRequest = func(w http.ResponseWriter, r *http.Request) (string, error) {
		return sessionID, err
	}
}
