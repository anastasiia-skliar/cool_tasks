package services

import (
	"github.com/Nastya-Kruglikova/cool_tasks/src/services/auth"
	"github.com/Nastya-Kruglikova/cool_tasks/src/services/common"
	"github.com/Nastya-Kruglikova/cool_tasks/src/services/events"
	"github.com/Nastya-Kruglikova/cool_tasks/src/services/flights"
	"github.com/Nastya-Kruglikova/cool_tasks/src/services/hotels"
	"github.com/Nastya-Kruglikova/cool_tasks/src/services/museums"
	"github.com/Nastya-Kruglikova/cool_tasks/src/services/tasksCRUD"
	"github.com/Nastya-Kruglikova/cool_tasks/src/services/trains"
	"github.com/Nastya-Kruglikova/cool_tasks/src/services/usersCRUD"
	"github.com/Nastya-Kruglikova/cool_tasks/src/services/welcome"
	"github.com/gorilla/mux"
	"net/http"
	"github.com/Nastya-Kruglikova/cool_tasks/src/services/trips"
	"github.com/Nastya-Kruglikova/cool_tasks/src/services/restaurants"
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
		http.MethodGet:  http.HandlerFunc(restaurants.Get),
		http.MethodPost: http.HandlerFunc(restaurants.SaveRestaurant),
	}))
	apiV1.Handle("/restaurants/trip/{id}", common.MethodHandler(map[string]http.Handler{
		http.MethodGet: http.HandlerFunc(restaurants.GetRestaurantFromTrip),
	}))

	apiV1.Handle("/users/trips/{id}", common.MethodHandler(map[string]http.Handler{
		http.MethodGet:  http.HandlerFunc(trips.GetTripIDsByUserIDHandler),
	}))
	apiV1.Handle("/trips", common.MethodHandler(map[string]http.Handler{
		http.MethodPost: http.HandlerFunc(trips.CreateTripHandler),
	}))
	apiV1.Handle("/trips/{id}", common.MethodHandler(map[string]http.Handler{
		http.MethodGet: http.HandlerFunc(trips.GetTripHandler),
	}))

	return router
}
