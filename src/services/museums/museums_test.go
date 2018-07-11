package museums_test

import (
	"bytes"
	"github.com/Nastya-Kruglikova/cool_tasks/src/models"
	"github.com/Nastya-Kruglikova/cool_tasks/src/services"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

var router = services.NewRouter()

type MuseumsTestCase struct {
	name             string
	url              string
	want             int
	mockedGetMuseums []models.Museum
}

func TestGetMuseumsByRequestHandler(t *testing.T) {
	tests := []MuseumsTestCase{
		{
			name:             "Get_Museums_200",
			url:              "/v1/museums",
			want:             200,
			mockedGetMuseums: []models.Museum{},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			models.MockedGetMuseums()
			rec := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodGet, tc.url, nil)

			router.ServeHTTP(rec, req)

			if rec.Code != tc.want {
				t.Errorf("Expected: %d , got %d", tc.want, rec.Code)
			}
		})
	}
}

func TestAddMuseumToTripHandler(t *testing.T) {
	tests := []MuseumsTestCase{
		{
			name: "Add_Museum_200",
			url:  "/v1/museums",
			want: 200,
		},
	}
	data := url.Values{}
	data.Add("museum", "00000000-0000-0000-0000-000000000001")
	data.Add("trip", "00000000-0000-0000-0000-000000000001")
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			models.MockedAddMuseum()
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

func TestGetMuseumByTripHandler(t *testing.T) {
	tests := []MuseumsTestCase{
		{
			name:             "Get_Museums_200",
			url:              "/v1/museums/trip/00000000-0000-0000-0000-000000000001",
			want:             200,
			mockedGetMuseums: []models.Museum{},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			models.MockedGetMuseumsByTrip()
			rec := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodGet, tc.url, nil)

			router.ServeHTTP(rec, req)

			if rec.Code != tc.want {
				t.Errorf("Expected: %d , got %d", tc.want, rec.Code)
			}
		})
	}
}
