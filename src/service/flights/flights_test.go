package flights_test

import (
	"bytes"
	"github.com/Nastya-Kruglikova/cool_tasks/src/model"
	"github.com/Nastya-Kruglikova/cool_tasks/src/service"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

var router = service.NewRouter()

type FlightsTestCase struct {
	name             string
	url              string
	want             int
	mockedGetFlights []model.Flight
	testDataId       string
	testDataFl       string
	mockedFlightsErr error
}

func TestGetByRequestHandler(t *testing.T) {
	tests := []FlightsTestCase{
		{
			name:             "Get_Flights_200",
			url:              "/v1/flights?departure_city=lviv&arrival_city=kyiv",
			want:             200,
			mockedGetFlights: []model.Flight{},
			mockedFlightsErr: nil,
		},
		{
			name:             "Get_Flights_400",
			url:              "/v1/flights?mock=890",
			want:             400,
			mockedGetFlights: []model.Flight{},
			mockedFlightsErr: http.ErrBodyNotAllowed,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			model.MockedGetData(tc.mockedGetFlights, tc.mockedFlightsErr)
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
			name:             "Add_To_Trip_201",
			url:              "/v1/flights",
			want:             201,
			testDataId:       "00000000-0000-0000-0000-000000000001",
			testDataFl:       "00000000-0000-0000-0000-000000000001",
			mockedFlightsErr: nil,
		},
		{
			name:             "Add_To_Trip_400",
			url:              "/v1/flights",
			want:             400,
			testDataId:       "00000000-0000-0000-0000-000000000001",
			testDataFl:       "asdas",
			mockedFlightsErr: nil,
		},
		{
			name:             "Add_Flights_400_2",
			url:              "/v1/flights",
			want:             400,
			testDataId:       "asdasd",
			testDataFl:       "00000000-0000-0000-0000-000000000001",
			mockedFlightsErr: nil,
		},
		{
			name:             "Add_Flights_400_3",
			url:              "/v1/flights",
			want:             400,
			testDataId:       "00000000-0000-0000-0000-000000000001",
			testDataFl:       "00000000-0000-0000-0000-000000000001",
			mockedFlightsErr: http.ErrLineTooLong,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			data := url.Values{}
			data.Add("flight_id", tc.testDataFl)
			data.Add("trip_id", tc.testDataId)

			model.MockedAddToTrip(tc.mockedFlightsErr)
			rec := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodPost, tc.url, bytes.NewBufferString(data.Encode()))
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")

			router.ServeHTTP(rec, req)

			if rec.Code != tc.want {
				t.Errorf("Expected: %d , got %d", tc.want, rec.Code)
			}
			data.Del("flight_id")
			data.Del("trip_id")
		})
	}
}

func TestGetByTripHandler(t *testing.T) {
	tests := []FlightsTestCase{
		{
			name:             "Get_flight_200",
			url:              "/v1/flights/trip/00000000-0000-0000-0000-000000000001",
			want:             200,
			mockedGetFlights: []model.Flight{},
			mockedFlightsErr: nil,
		},
		{
			name:             "Get_flight_400",
			url:              "/v1/flights/trip/asdas",
			want:             400,
			mockedGetFlights: []model.Flight{},
			mockedFlightsErr: nil,
		},
		{
			name:             "Get_Flights_404",
			url:              "/v1/flights/trip/00000000-0000-0000-0000-000000000009",
			want:             404,
			mockedGetFlights: []model.Flight{},
			mockedFlightsErr: http.ErrLineTooLong,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			model.MockedGetByTrip(tc.mockedGetFlights, tc.mockedFlightsErr)
			rec := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodGet, tc.url, nil)

			router.ServeHTTP(rec, req)

			if rec.Code != tc.want {
				t.Errorf("Expected: %d , got %d", tc.want, rec.Code)
			}
		})
	}
}
