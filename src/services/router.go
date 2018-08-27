package services

import (
	"github.com/Nastya-Kruglikova/cool_tasks/src/services/auth"
	"github.com/Nastya-Kruglikova/cool_tasks/src/services/common"

	"github.com/Nastya-Kruglikova/cool_tasks/src/services/events"
	"github.com/Nastya-Kruglikova/cool_tasks/src/services/flights"
	"github.com/Nastya-Kruglikova/cool_tasks/src/services/hotels"
	"github.com/Nastya-Kruglikova/cool_tasks/src/services/museums"
	"github.com/Nastya-Kruglikova/cool_tasks/src/services/restaurants"
	"github.com/Nastya-Kruglikova/cool_tasks/src/services/tasks"
	"github.com/Nastya-Kruglikova/cool_tasks/src/services/trains"
	"github.com/Nastya-Kruglikova/cool_tasks/src/services/trips"
	"github.com/Nastya-Kruglikova/cool_tasks/src/services/users"
	"github.com/Nastya-Kruglikova/cool_tasks/src/services/welcome"
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
		http.MethodGet:  http.HandlerFunc(users.GetUsersHandler),
		http.MethodPost: http.HandlerFunc(users.AddUserHandler),
	}))
	apiV1.Handle("/users/{id}", common.MethodHandler(map[string]http.Handler{
		http.MethodGet:    http.HandlerFunc(users.GetUserHandler),
		http.MethodDelete: http.HandlerFunc(users.DeleteUserHandler),
	}))
	apiV1.Handle("/users/tasks/{id}", common.MethodHandler(map[string]http.Handler{
		http.MethodGet: http.HandlerFunc(tasks.GetUserTasksHandler),
	}))

	apiV1.Handle("/tasks", common.MethodHandler(map[string]http.Handler{
		http.MethodGet:  http.HandlerFunc(tasks.GetTasksHandler),
		http.MethodPost: http.HandlerFunc(tasks.AddTaskHandler),
	}))
	apiV1.Handle("/tasks/{id}", common.MethodHandler(map[string]http.Handler{
		http.MethodGet:    http.HandlerFunc(tasks.GetTaskHandler),
		http.MethodDelete: http.HandlerFunc(tasks.DeleteTaskHandler),
	}))
	apiV1.Handle("/restaurants", common.MethodHandler(map[string]http.Handler{
		http.MethodPost: http.HandlerFunc(restaurants.AddRestaurantToTripHandler),
		http.MethodGet:  http.HandlerFunc(restaurants.GetRestaurantHandler),
	}))
	apiV1.Handle("/restaurants/trip/{id}", common.MethodHandler(map[string]http.Handler{
		http.MethodGet: http.HandlerFunc(restaurants.GetRestaurantFromTrip),
	}))
	apiV1.Handle("/events", common.MethodHandler(map[string]http.Handler{
		http.MethodGet:  http.HandlerFunc(events.GetEventsHandler),
		http.MethodPost: http.HandlerFunc(events.AddEventToTripHandler),
	}))
	apiV1.Handle("/events/trip/{id}", common.MethodHandler(map[string]http.Handler{
		http.MethodGet: http.HandlerFunc(events.GetEventsByTripHandler),
	}))

	apiV1.Handle("/flights", common.MethodHandler(map[string]http.Handler{
		http.MethodGet:  http.HandlerFunc(flights.GetFlightsHandler),
		http.MethodPost: http.HandlerFunc(flights.AddFlightToTripHandler),
	}))
	apiV1.Handle("/flights/trip/{id}", common.MethodHandler(map[string]http.Handler{
		http.MethodGet: http.HandlerFunc(flights.GetFlightsByTripHandler),
	}))
	apiV1.Handle("/museums", common.MethodHandler(map[string]http.Handler{
		http.MethodGet:  http.HandlerFunc(museums.GetMuseumsHandler),
		http.MethodPost: http.HandlerFunc(museums.AddMuseumToTripHandler),
	}))
	apiV1.Handle("/museums/trip/{id}", common.MethodHandler(map[string]http.Handler{
		http.MethodGet: http.HandlerFunc(museums.GetMuseumsByTripHandler),
	}))

	apiV1.Handle("/trains", common.MethodHandler(map[string]http.Handler{
		http.MethodGet:  http.HandlerFunc(trains.GetTrainsHandler),
		http.MethodPost: http.HandlerFunc(trains.AddTrainToTripHandler),
	}))
	apiV1.Handle("/trains/trip/{id}", common.MethodHandler(map[string]http.Handler{
		http.MethodGet: http.HandlerFunc(trains.GetTrainsFromTripHandler),
	}))
	apiV1.Handle("/hotels", common.MethodHandler(map[string]http.Handler{
		http.MethodGet:  http.HandlerFunc(hotels.GetHotelsHandler),
		http.MethodPost: http.HandlerFunc(hotels.AddHotelToTripHandler),
	}))
	apiV1.Handle("/hotels/trip/{id}", common.MethodHandler(map[string]http.Handler{
		http.MethodGet: http.HandlerFunc(hotels.GetHotelsByTripHandler),
	}))

	apiV1.Handle("/restaurants", common.MethodHandler(map[string]http.Handler{
		http.MethodGet:  http.HandlerFunc(restaurants.GetRestaurantHandler),
		http.MethodPost: http.HandlerFunc(restaurants.AddRestaurantToTripHandler),
	}))
	apiV1.Handle("/restaurants/trip/{id}", common.MethodHandler(map[string]http.Handler{
		http.MethodGet: http.HandlerFunc(restaurants.GetRestaurantFromTrip),
	}))

	apiV1.Handle("/users/trips/{id}", common.MethodHandler(map[string]http.Handler{
		http.MethodGet: http.HandlerFunc(trips.GetTripIDsByUserIDHandler),
	}))
	apiV1.Handle("/trips", common.MethodHandler(map[string]http.Handler{
		http.MethodPost: http.HandlerFunc(trips.AddTripHandler),
	}))
	apiV1.Handle("/trips/{id}", common.MethodHandler(map[string]http.Handler{
		http.MethodGet: http.HandlerFunc(trips.GetTripHandler),
	}))

	return router
}
