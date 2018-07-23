package museums_test

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

type MuseumsTestCase struct {
	name             string
	url              string
	want             int
	mockedGetMuseums []models.Museum
	testDataId       string
	testDataMu       string
	mock             func()
}

func TestGetMuseumsByRequestHandler(t *testing.T) {
	tests := []MuseumsTestCase{
		{
			name:             "Get_Museums_200",
			url:              "/v1/museums",
			want:             200,
			mockedGetMuseums: []models.Museum{},
			mock: func() {

			},
		},
		{
			name:             "Get_Museums_404",
			url:              "/v1/museums?mock=890",
			want:             404,
			mockedGetMuseums: []models.Museum{},
			mock: func() {
				var err = http.ErrBodyNotAllowed
				models.GetMuseums = func(values url.Values) ([]models.Museum, error) {
					return []models.Museum{}, err
				}
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			models.MockedGetMuseums()
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

func TestAddMuseumToTripHandler(t *testing.T) {
	tests := []MuseumsTestCase{
		{
			name:       "Add_Museum_200",
			url:        "/v1/museums",
			want:       200,
			testDataId: "00000000-0000-0000-0000-000000000001",
			testDataMu: "00000000-0000-0000-0000-000000000001",
			mock: func() {

			},
		},
		{
			name:       "Add_Museums_400",
			url:        "/v1/museums",
			want:       400,
			testDataId: "00000000-0000-0000-0000-000000000001",
			testDataMu: "asdsad",
			mock: func() {

			},
		},
		{
			name:       "Add_Museums_400_2",
			url:        "/v1/museums",
			want:       400,
			testDataId: "asdasd",
			testDataMu: "00000000-0000-0000-0000-000000000001",
			mock: func() {

			},
		},
		{
			name:       "Add_Museums_400_3",
			url:        "/v1/museums",
			want:       400,
			testDataId: "00000000-0000-0000-0000-000000000001",
			testDataMu: "00000000-0000-0000-0000-000000000001",
			mock: func() {
				var err = error(new(http.ProtocolError))
				models.AddMuseumToTrip = func(museum_id uuid.UUID, trip_id uuid.UUID) error {
					return err
				}
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			data := url.Values{}
			data.Add("museum", tc.testDataMu)
			data.Add("trip", tc.testDataId)
			models.MockedAddMuseum()
			tc.mock()
			rec := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodPost, tc.url, bytes.NewBufferString(data.Encode()))
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")

			router.ServeHTTP(rec, req)

			if rec.Code != tc.want {
				t.Errorf("Expected: %d , got %d", tc.want, rec.Code)
			}
			data.Del("museum")
			data.Del("trip")
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
			mock: func() {

			},
		},
		{
			name:             "Get_Museums_200",
			url:              "/v1/museums/trip/sadsad",
			want:             400,
			mockedGetMuseums: []models.Museum{},
			mock: func() {

			},
		},
		{
			name:             "Get_Museums_404",
			url:              "/v1/museums/trip/00000000-0000-0000-0000-000000000009",
			want:             404,
			mockedGetMuseums: []models.Museum{},
			mock: func() {
				var err = http.ErrLineTooLong
				models.GetMuseumsByTrip = func(trip_id uuid.UUID) ([]models.Museum, error) {
					return []models.Museum{}, err
				}
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			models.MockedGetMuseumsByTrip()
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
