package events_test

import (
	"testing"
	"net/http/httptest"
	"net/http"
	"github.com/Nastya-Kruglikova/cool_tasks/src/services"
	"github.com/Nastya-Kruglikova/cool_tasks/src/models"
	"net/url"
	"bytes"
)


var router = services.NewRouter()

type EventsTestCase struct {
	name             string
	url              string
	want             int
	mockedGetEvents []models.Event
}

func TestGetByRequestHandler(t *testing.T) {
	tests := []EventsTestCase{
		{
			name:             "Get_Events_200",
			url:              "/v1/events",
			want:             200,
			mockedGetEvents: []models.Event{},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			models.MockedGetEvents()
			rec := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodGet, tc.url, nil)

			router.ServeHTTP(rec, req)

			if rec.Code != tc.want {
				t.Errorf("Expected: %d , got %d", tc.want, rec.Code)
			}
		})
	}
}


func TestAddToTripHandler(t *testing.T) {
	tests := []EventsTestCase{
		{
			name: "Add_Events_200",
			url:  "/v1/events",
			want: 200,
		},
	}
	data := url.Values{}
	data.Add("event_id", "00000000-0000-0000-0000-000000000001")
	data.Add("trip_id", "00000000-0000-0000-0000-000000000001")
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			models.MockedAddEventToTrip()
			rec := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodPost, tc.url, bytes.NewBufferString(data.Encode()))
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")

			router.ServeHTTP(rec, req)

			if rec.Code != tc.want {
				t.Errorf("Expected: %d , got %d", tc.want, rec.Code)
			}
		})
	}
}

func TestGetByTripHandler(t *testing.T) {
	tests := []EventsTestCase{
		{
			name:             "Get_Events_200",
			url:              "/v1/events/trip/00000000-0000-0000-0000-000000000001",
			want:             200,
			mockedGetEvents: []models.Event{},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			models.MockedGetEventsByTrip()
			rec := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodGet, tc.url, nil)

			router.ServeHTTP(rec, req)

			if rec.Code != tc.want {
				t.Errorf("Expected: %d , got %d", tc.want, rec.Code)
			}
		})
	}
}
