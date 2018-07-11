package hotels_test

import (
	"testing"
	"github.com/Nastya-Kruglikova/cool_tasks/src/models"
	"net/http/httptest"
	"net/http"
	"github.com/Nastya-Kruglikova/cool_tasks/src/services"
	"net/url"
	"bytes"
)

var router = services.NewRouter()

type HotelsTestCase struct {
	name             string
	url              string
	want             int
	mockedGetHotels []models.Hotel
}

func TestGetHotels(t *testing.T) {
	tests := []HotelsTestCase{
		{
			name:             "Get_Hotels_200",
			url:              "/v1/hotels",
			want:             200,
			mockedGetHotels: []models.Hotel{},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			models.GetHotelsMocked()
			rec := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodGet, tc.url, nil)

			router.ServeHTTP(rec, req)

			if rec.Code != tc.want {
				t.Errorf("Expected: %d , got %d", tc.want, rec.Code)
			}
		})
	}
}


func TestAddHotel(t *testing.T) {
	tests := []HotelsTestCase{
		{
			name: "Add_Hotels_200",
			url:  "/v1/hotels",
			want: 200,
		},
	}
	data := url.Values{}
	data.Add("hotel", "00000000-0000-0000-0000-000000000001")
	data.Add("trip", "00000000-0000-0000-0000-000000000001")
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			models.AddHotelMocked()
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

func TestGetHotelByTripHandler(t *testing.T) {
	tests := []HotelsTestCase{
		{
			name:             "Get_Hotels_200",
			url:              "/v1/hotels/trip/00000000-0000-0000-0000-000000000001",
			want:             200,
			mockedGetHotels: []models.Hotel{},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			models.GetHotelByTripIdMocked()
			rec := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodGet, tc.url, nil)

			router.ServeHTTP(rec, req)

			if rec.Code != tc.want {
				t.Errorf("Expected: %d , got %d", tc.want, rec.Code)
			}
		})
	}
}
