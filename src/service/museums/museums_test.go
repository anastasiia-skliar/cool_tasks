package museums_test

import (
	"bytes"
	"encoding/json"
	"github.com/Nastya-Kruglikova/cool_tasks/src/model"
	"github.com/Nastya-Kruglikova/cool_tasks/src/service"
	"github.com/Nastya-Kruglikova/cool_tasks/src/service/museums"
	"net/http"
	"net/http/httptest"
	"testing"
)

var router = service.NewRouter()

type MuseumsTestCase struct {
	name             string
	url              string
	want             int
	mockedGetMuseums []model.Museum
	testDataId       string
	testDataMu       string
	mockedMuseumErr  error
}

func TestGetMuseumsByRequestHandler(t *testing.T) {
	tests := []MuseumsTestCase{
		{
			name:             "Get_Museums_200",
			url:              "/v1/museums",
			want:             200,
			mockedGetMuseums: []model.Museum{},
			mockedMuseumErr:  nil,
		},
		{
			name:             "Get_Museums_404",
			url:              "/v1/museums?mock=890",
			want:             404,
			mockedGetMuseums: []model.Museum{},
			mockedMuseumErr:  http.ErrBodyNotAllowed,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			model.MockedGetData(tc.mockedGetMuseums, tc.mockedMuseumErr)
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
			name:            "Add_Museum_201",
			url:             "/v1/museums",
			want:            201,
			testDataId:      "00000000-0000-0000-0000-000000000001",
			testDataMu:      "00000000-0000-0000-0000-000000000001",
			mockedMuseumErr: nil,
		},
		{
			name:            "Add_Museums_400",
			url:             "/v1/museums",
			want:            400,
			testDataId:      "00000000-0000-0000-0000-000000000001",
			testDataMu:      "asdsad",
			mockedMuseumErr: nil,
		},
		{
			name:            "Add_Museums_400_2",
			url:             "/v1/museums",
			want:            400,
			testDataId:      "asdasd",
			testDataMu:      "00000000-0000-0000-0000-000000000001",
			mockedMuseumErr: nil,
		},
		{
			name:            "Add_Museums_400_3",
			url:             "/v1/museums",
			want:            400,
			testDataId:      "00000000-0000-0000-0000-000000000001",
			testDataMu:      "00000000-0000-0000-0000-000000000001",
			mockedMuseumErr: error(new(http.ProtocolError)),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			var data museums.TripMuseum
			data.MuseumID = tc.testDataMu
			data.TripID = tc.testDataId
			body, _ := json.Marshal(data)

			model.MockedAddToTrip(tc.mockedMuseumErr)
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

func TestGetMuseumByTripHandler(t *testing.T) {
	tests := []MuseumsTestCase{
		{
			name:             "Get_Museums_200",
			url:              "/v1/museums/trip/00000000-0000-0000-0000-000000000001",
			want:             200,
			mockedGetMuseums: []model.Museum{},
			mockedMuseumErr:  nil,
		},
		{
			name:             "Get_Museums_200",
			url:              "/v1/museums/trip/sadsad",
			want:             400,
			mockedGetMuseums: []model.Museum{},
			mockedMuseumErr:  nil,
		},
		{
			name:             "Get_Museums_404",
			url:              "/v1/museums/trip/00000000-0000-0000-0000-000000000009",
			want:             404,
			mockedGetMuseums: []model.Museum{},
			mockedMuseumErr:  http.ErrLineTooLong,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			model.MockedGetByTrip(tc.mockedGetMuseums, tc.mockedMuseumErr)
			rec := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodGet, tc.url, nil)

			router.ServeHTTP(rec, req)

			if rec.Code != tc.want {
				t.Errorf("Expected: %d , got %d", tc.want, rec.Code)
			}
		})
	}
}
