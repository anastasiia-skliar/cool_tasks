package auth

import (
	"net/http"
	"net/http/httptest"
)

func MockedLogin(recorder *httptest.ResponseRecorder, request *http.Request) {
	Login = func(w http.ResponseWriter, r *http.Request) {

	}
}
func MockedLogout(recorder *httptest.ResponseRecorder, request *http.Request) {
	Logout = func(w http.ResponseWriter, r *http.Request) {

	}
}
