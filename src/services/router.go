package services

import (
	"net/http"
	"github.com/Nastya-Kruglikova/cool_tasks/src/services/usersCRUD"
	"github.com/Nastya-Kruglikova/cool_tasks/src/services/common"
	"github.com/Nastya-Kruglikova/cool_tasks/src/services/welcome"
	"github.com/gorilla/mux"
)

// NewRouter creates a router for URL-to-service mapping
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	apiV1 := router.PathPrefix("/v1/").Subrouter()

	apiV1.Handle("/hello-world", common.MethodHandler(map[string]http.Handler{
		http.MethodGet: http.HandlerFunc(welcome.GetWelcomeHandler),
	}))
	apiV1.Handle("/users", common.MethodHandler(map[string]http.Handler{
		http.MethodGet:  http.HandlerFunc(usersCRUD.GetUsers),
		http.MethodPost: http.HandlerFunc(usersCRUD.AddUser),
	}))
	apiV1.Handle("/users/{id}", common.MethodHandler(map[string]http.Handler{
		http.MethodGet:    http.HandlerFunc(usersCRUD.GetUserByID),
		http.MethodDelete: http.HandlerFunc(usersCRUD.DeleteUser),
	}))
	return router
}
