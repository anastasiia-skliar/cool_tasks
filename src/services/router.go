package services

import (
	"net/http"

	"github.com/Nastya-Kruglikova/cool_tasks/src/services/common"
	"github.com/Nastya-Kruglikova/cool_tasks/src/services/welcome"
	"github.com/Nastya-Kruglikova/cool_tasks/src/services/tasksCRUD"
)

// NewRouter creates a router for URL-to-service mapping
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	apiV1 := router.PathPrefix("/v1/").Subrouter()

	apiV1.Handle("/hello-world", common.MethodHandler(map[string]http.Handler{
		http.MethodGet: http.HandlerFunc(welcome.GetWelcomeHandler),
	}))


	apiV1.Handle("/tasks", common.MethodHandler(map[string]http.Handler{
		http.MethodGet: http.HandlerFunc(tasksCRUD.GetTasks),
	}))
	apiV1.Handle("/tasks/{id}", common.MethodHandler(map[string]http.Handler{
		http.MethodGet: http.HandlerFunc(tasksCRUD.GetTasksByID),
	}))
	apiV1.Handle("/tasks", common.MethodHandler(map[string]http.Handler{
		http.MethodPost: http.HandlerFunc(tasksCRUD.AddTasks),
	}))
	apiV1.Handle("/tasks/{id}", common.MethodHandler(map[string]http.Handler{
		http.MethodDelete: http.HandlerFunc(tasksCRUD.DeleteTasks),
	}))

	return router
}
