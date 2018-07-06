package CRUD_test

import (
	"testing"
	"net/http/httptest"
	"net/http"
	"github.com/Nastya-Kruglikova/cool_tasks/src/services"
	"github.com/Nastya-Kruglikova/cool_tasks/src/flights/models"
	"net/url"
	"bytes"
)

var router = services.NewRouter()

type FlightsTestCase struct {
	name string
	url  string
	want int
}

func TestGetByRequestHandler(t *testing.T) {
	tests := []FlightsTestCase{
		{
			name: "Get_Flights_200",
			url:  "/v1/flights?departure_city=lviv&arrival_city=kyiv",
			want: 200,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			models.MockedGetByRequest()
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
	tests := []FlightsTestCase{
		{
			name: "Add_To_Trip_200",
			url:  "/v1/flights",
			want: 200,
		},
	}
	data := url.Values{}
	data.Add("flight_id", "00000000-0000-0000-0000-000000000001")
	data.Add("trip_id", "00000000-0000-0000-0000-000000000001")

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			models.MockedAddToTrip()
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
	tests := []FlightsTestCase{
		{
			name: "Get_flight_200",
			url:  "/v1/flights/trip/00000000-0000-0000-0000-000000000001",
			want: 200,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			models.MockedGetByTrip()
			rec := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodGet, tc.url, nil)

			router.ServeHTTP(rec, req)

			if rec.Code != tc.want {
				t.Errorf("Expected: %d , got %d", tc.want, rec.Code)
			}
		})
	}
}
