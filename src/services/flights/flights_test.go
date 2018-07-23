package flights_test

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

type FlightsTestCase struct {
	name             string
	url              string
	want             int
	mockedGetFlights []models.Flight
	testDataId       string
	testDataFl       string
	mock             func()
}

func TestGetByRequestHandler(t *testing.T) {
	tests := []FlightsTestCase{
		{
			name: "Get_Flights_200",
			url:  "/v1/flights?departure_city=lviv&arrival_city=kyiv",
			want: 200,
			mock: func() {

			},
		},
		{
			name:             "Get_Flights_400",
			url:              "/v1/flights?mock=890",
			want:             400,
			mockedGetFlights: []models.Flight{},
			mock: func() {
				var err = http.ErrBodyNotAllowed
				models.GetFlightsByRequest = func(values url.Values) ([]models.Flight, error) {
					return []models.Flight{}, err
				}
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			models.MockedGetByRequest()
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

func TestAddToTripHandler(t *testing.T) {
	tests := []FlightsTestCase{
		{
			name:       "Add_To_Trip_200",
			url:        "/v1/flights",
			want:       200,
			testDataId: "00000000-0000-0000-0000-000000000001",
			testDataFl: "00000000-0000-0000-0000-000000000001",
			mock: func() {

			},
		},
		{
			name:       "Add_To_Trip_400",
			url:        "/v1/flights",
			want:       400,
			testDataId: "00000000-0000-0000-0000-000000000001",
			testDataFl: "asdas",
			mock: func() {

			},
		},
		{
			name:       "Add_Flights_400_2",
			url:        "/v1/flights",
			want:       400,
			testDataId: "asdasd",
			testDataFl: "00000000-0000-0000-0000-000000000001",
			mock: func() {

			},
		},
		{
			name:             "Add_Flights_400_3",
			url:              "/v1/flights",
			want:             400,
			testDataId:       "00000000-0000-0000-0000-000000000001",
			testDataFl:       "00000000-0000-0000-0000-000000000001",
			mockedGetFlights: []models.Flight{},
			mock: func() {
				var err = http.ErrLineTooLong
				models.AddFlightToTrip = func(flightID uuid.UUID, trip_id uuid.UUID) error {
					return err
				}
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			data := url.Values{}
			data.Add("flight_id", tc.testDataFl)
			data.Add("trip_id", tc.testDataId)

			models.MockedAddToTrip()
			tc.mock()
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
			mockedGetFlights: []models.Flight{},
			mock: func() {

			},
		},
		{
			name:             "Get_flight_400",
			url:              "/v1/flights/trip/asdas",
			want:             400,
			mockedGetFlights: []models.Flight{},
			mock: func() {

			},
		},
		{
			name:             "Get_Flights_404",
			url:              "/v1/flights/trip/00000000-0000-0000-0000-000000000009",
			want:             404,
			mockedGetFlights: []models.Flight{},
			mock: func() {
				var err = http.ErrLineTooLong
				models.GetFlightsByTrip = func(trip_id uuid.UUID) ([]models.Flight, error) {
					return []models.Flight{}, err
				}
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			models.MockedGetByTrip()
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
