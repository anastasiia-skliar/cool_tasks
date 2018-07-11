package services

import (
	"github.com/Nastya-Kruglikova/cool_tasks/src/services/auth"
	"github.com/Nastya-Kruglikova/cool_tasks/src/services/common"
	"github.com/Nastya-Kruglikova/cool_tasks/src/services/tasksCRUD"
	"github.com/Nastya-Kruglikova/cool_tasks/src/services/usersCRUD"
	"github.com/Nastya-Kruglikova/cool_tasks/src/services/welcome"
	"github.com/Nastya-Kruglikova/cool_tasks/src/services/events"
	"github.com/gorilla/mux"
	"net/http"
)

// NewRouter creates a router for URL-to-service mapping
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	apiV1 := router.PathPrefix("/v1/").Subrouter()

	apiV1.Handle("/hello-world", common.MethodHandler(map[string]http.Handler{
		http.MethodGet: http.HandlerFunc(welcome.GetWelcomeHandler),
	}))

	apiV1.Handle("/login", common.MethodHandler(map[string]http.Handler{
		http.MethodPost: http.HandlerFunc(auth.Login),
	}))
	apiV1.Handle("/logout", common.MethodHandler(map[string]http.Handler{
		http.MethodPost: http.HandlerFunc(auth.Logout),
	}))

	apiV1.Handle("/users", common.MethodHandler(map[string]http.Handler{
		http.MethodGet:  http.HandlerFunc(usersCRUD.GetUsers),
		http.MethodPost: http.HandlerFunc(usersCRUD.CreateUser),
	}))
	apiV1.Handle("/users/{id}", common.MethodHandler(map[string]http.Handler{
		http.MethodGet:    http.HandlerFunc(usersCRUD.GetUserByID),
		http.MethodDelete: http.HandlerFunc(usersCRUD.DeleteUser),
	}))
	apiV1.Handle("/users/tasks/{id}", common.MethodHandler(map[string]http.Handler{
		http.MethodGet: http.HandlerFunc(tasksCRUD.GetUserTasks),
	}))

	apiV1.Handle("/tasks", common.MethodHandler(map[string]http.Handler{
		http.MethodGet:  http.HandlerFunc(tasksCRUD.GetTasks),
		http.MethodPost: http.HandlerFunc(tasksCRUD.CreateTask),
	}))
	apiV1.Handle("/tasks/{id}", common.MethodHandler(map[string]http.Handler{
		http.MethodGet:    http.HandlerFunc(tasksCRUD.GetTasksByID),
		http.MethodDelete: http.HandlerFunc(tasksCRUD.DeleteTasks),
	}))
	
	apiV1.Handle("/events", common.MethodHandler(map[string]http.Handler{
		http.MethodGet:  http.HandlerFunc(events.GetByRequestHandler),
		http.MethodPost: http.HandlerFunc(events.AddToTripHandler),
	}))
	apiV1.Handle("/events/trip/{id}", common.MethodHandler(map[string]http.Handler{
		http.MethodGet: http.HandlerFunc(events.GetByTripHandler),
	}))

	return router
}
