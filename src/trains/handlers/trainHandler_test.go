package handlers_test

import (
	"bytes"
	"github.com/Nastya-Kruglikova/cool_tasks/src/services"
	"github.com/Nastya-Kruglikova/cool_tasks/src/trains/models"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

var router = services.NewRouter()

type TrainsTestCase struct {
	name string
	url  string
	want int
}

func TestGetTrains(t *testing.T) {
	tests := []TrainsTestCase{
		{
			name: "Get_Trains_200",
			url:  "/v1/trains?departure_city=lviv&arrival_city=kyiv&price=300uah",
			want: 200,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			models.MockedGetTrains()
			rec := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodGet, tc.url, nil)

			router.ServeHTTP(rec, req)

			if rec.Code != tc.want {
				t.Errorf("Expected: %d , got %d", tc.want, rec.Code)
			}
		})
	}
}

func TestSaveTrain(t *testing.T) {
	tests := []TrainsTestCase{
		{
			name: "Add_To_Trip_200",
			url:  "/v1/trains",
			want: 200,
		},
	}
	data := url.Values{}
	data.Add("train_id", "00000000-0000-0000-0000-000000000001")
	data.Add("trip_id", "00000000-0000-0000-0000-000000000001")

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			models.MockedSaveTrain()
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

func TestGetFromTrip(t *testing.T) {
	tests := []TrainsTestCase{
		{
			name: "Get_train_200",
			url:  "/v1/trains/trip/00000000-0000-0000-0000-000000000001",
			want: 200,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			models.MockedGetFromTrip()
			rec := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodGet, tc.url, nil)

			router.ServeHTTP(rec, req)

			if rec.Code != tc.want {
				t.Errorf("Expected: %d , got %d", tc.want, rec.Code)
			}
		})
	}
}
