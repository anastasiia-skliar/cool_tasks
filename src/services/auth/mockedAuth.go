package auth

import (
	"net/http"
	"net/http/httptest"
)

//MockedLogin is mocked Login func
func MockedLogin(recorder *httptest.ResponseRecorder, request *http.Request) {
	Login = func(w http.ResponseWriter, r *http.Request) {

	}
}

//MockedLogout is mocked Logout func
func MockedLogout(recorder *httptest.ResponseRecorder, request *http.Request) {
	Logout = func(w http.ResponseWriter, r *http.Request) {

	}
}
