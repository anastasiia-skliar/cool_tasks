package hotels_test

import (
	"bytes"
	"github.com/Nastya-Kruglikova/cool_tasks/src/models"
	"github.com/Nastya-Kruglikova/cool_tasks/src/services"
	"github.com/satori/go.uuid"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

var router = services.NewRouter()

type HotelsTestCase struct {
	name            string
	url             string
	want            int
	mockedGetHotels []models.Hotel
	testDataId      string
	testDataHo      string
	mock            func()
}

func TestGetHotels(t *testing.T) {
	tests := []HotelsTestCase{
		{
			name:            "Get_Hotels_200",
			url:             "/v1/hotels",
			want:            200,
			mockedGetHotels: []models.Hotel{},
			mock: func() {

			},
		},
		{
			name:            "Get_Hotels_404",
			url:             "/v1/hotels?mock=890",
			want:            404,
			mockedGetHotels: []models.Hotel{},
			mock: func() {
				var err = http.ErrBodyNotAllowed
				models.GetHotels = func(values url.Values) ([]models.Hotel, error) {
					return []models.Hotel{}, err
				}
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			models.GetHotelsMocked()
			rec := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodGet, tc.url, nil)
			tc.mock()
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
			name:       "Add_Hotels_200",
			url:        "/v1/hotels",
			want:       200,
			testDataId: "00000000-0000-0000-0000-000000000001",
			testDataHo: "00000000-0000-0000-0000-000000000001",
			mock: func() {

			},
		},
		{
			name:       "Add_Hotels_400",
			url:        "/v1/hotels",
			want:       400,
			testDataId: "00000000-0000-0000-0000-000000000001",
			testDataHo: "asdsad",
			mock: func() {

			},
		},
		{
			name:       "Add_Hotels_400_2",
			url:        "/v1/hotels",
			want:       400,
			testDataId: "asdasd",
			testDataHo: "00000000-0000-0000-0000-000000000001",
			mock: func() {

			},
		},
		{
			name:       "Add_Hotels_400_3",
			url:        "/v1/hotels",
			want:       400,
			testDataId: "00000000-0000-0000-0000-000000000001",
			testDataHo: "00000000-0000-0000-0000-000000000001",
			mock: func() {
				var err = error(new(http.ProtocolError))
				models.AddHotelToTrip = func(hotel_id uuid.UUID, trip_id uuid.UUID) error {
					return err
				}
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			data := url.Values{}
			data.Add("hotel_id", tc.testDataHo)
			data.Add("trip_id", tc.testDataId)
			models.AddHotelMocked()
			tc.mock()
			rec := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodPost, tc.url, bytes.NewBufferString(data.Encode()))
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
			router.ServeHTTP(rec, req)
			if rec.Code != tc.want {
				t.Errorf("Expected: %d , got %d", tc.want, rec.Code)
			}
			data.Del("hotel_id")
			data.Del("trip_id")
		})
	}
}

func TestGetHotelByTripHandler(t *testing.T) {
	tests := []HotelsTestCase{
		{
			name: "Get_Hotels_200",
			url:  "/v1/hotels/trip/00000000-0000-0000-0000-000000000001",
			want: 200,
			mock: func() {

			},
		},
		{
			name: "Get_Hotels_400",
			url:  "/v1/hotels/trip/sadsad",
			want: 400,
			mock: func() {

			},
		},
		{
			name:            "Get_Hotels_404",
			url:             "/v1/hotels/trip/00000000-0000-0000-0000-000000000009",
			want:            404,
			mockedGetHotels: []models.Hotel{},
			mock: func() {
				var err = http.ErrLineTooLong
				models.GetHotelsByTrip = func(trip_id uuid.UUID) ([]models.Hotel, error) {
					return []models.Hotel{}, err
				}
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			models.GetHotelByTripIDMocked()
			tc.mock()
			rec := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodGet, tc.url, nil)

			router.ServeHTTP(rec, req)

			if rec.Code != tc.want {
				t.Errorf("Expected: %d , got %d", tc.want, rec.Code)
			}
		})
	}
}
