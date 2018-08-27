package hotels_test

import (
	"bytes"
	"encoding/json"
	"github.com/Nastya-Kruglikova/cool_tasks/src/model"
	"github.com/Nastya-Kruglikova/cool_tasks/src/service"
	"github.com/Nastya-Kruglikova/cool_tasks/src/service/hotels"
	"net/http"
	"net/http/httptest"
	"testing"
)

var router = service.NewRouter()

type HotelsTestCase struct {
	name            string
	url             string
	want            int
	mockedGetHotels []model.Hotel
	testDataId      string
	testDataHo      string
	mockedHotelErr  error
}

func TestGetHotels(t *testing.T) {
	tests := []HotelsTestCase{
		{
			name:            "Get_Hotels_200",
			url:             "/v1/hotels",
			want:            200,
			mockedGetHotels: []model.Hotel{},
			mockedHotelErr:  nil,
		},
		{
			name:            "Get_Hotels_404",
			url:             "/v1/hotels?mock=890",
			want:            404,
			mockedGetHotels: []model.Hotel{},
			mockedHotelErr:  http.ErrBodyNotAllowed,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			model.GetHotelsMocked(tc.mockedGetHotels, tc.mockedHotelErr)
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
			name:           "Add_Hotels_201",
			url:            "/v1/hotels",
			want:           201,
			testDataId:     "00000000-0000-0000-0000-000000000001",
			testDataHo:     "00000000-0000-0000-0000-000000000001",
			mockedHotelErr: nil,
		},
		{
			name:           "Add_Hotels_400",
			url:            "/v1/hotels",
			want:           400,
			testDataId:     "00000000-0000-0000-0000-000000000001",
			testDataHo:     "asdsad",
			mockedHotelErr: nil,
		},
		{
			name:           "Add_Hotels_400_2",
			url:            "/v1/hotels",
			want:           400,
			testDataId:     "asdasd",
			testDataHo:     "00000000-0000-0000-0000-000000000001",
			mockedHotelErr: nil,
		},
		{
			name:           "Add_Hotels_400_3",
			url:            "/v1/hotels",
			want:           400,
			testDataId:     "00000000-0000-0000-0000-000000000001",
			testDataHo:     "00000000-0000-0000-0000-000000000001",
			mockedHotelErr: error(new(http.ProtocolError)),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			var data hotels.TripHotel
			data.HotelID = tc.testDataHo
			data.TripID = tc.testDataId
			body, _ := json.Marshal(data)

			model.AddHotelMocked(tc.mockedHotelErr)
			rec := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodPost, tc.url, bytes.NewReader(body))
			req.Header.Set("Content-Type", "application/json")
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
			name:            "Get_Hotels_200",
			url:             "/v1/hotels/trip/00000000-0000-0000-0000-000000000001",
			want:            200,
			mockedGetHotels: []model.Hotel{},
			mockedHotelErr:  nil,
		},
		{
			name:            "Get_Hotels_400",
			url:             "/v1/hotels/trip/sadsad",
			want:            400,
			mockedGetHotels: []model.Hotel{},
			mockedHotelErr:  nil,
		},
		{
			name:            "Get_Hotels_404",
			url:             "/v1/hotels/trip/00000000-0000-0000-0000-000000000009",
			want:            404,
			mockedGetHotels: []model.Hotel{},
			mockedHotelErr:  http.ErrLineTooLong,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			model.GetHotelByTripIDMocked(tc.mockedGetHotels, tc.mockedHotelErr)
			rec := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodGet, tc.url, nil)

			router.ServeHTTP(rec, req)

			if rec.Code != tc.want {
				t.Errorf("Expected: %d , got %d", tc.want, rec.Code)
			}
		})
	}
}
