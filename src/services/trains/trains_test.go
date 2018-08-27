package trains_test

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

type TrainsTestCase struct {
	name             string
	url              string
	want             int
	mockedGetTrains  []models.Train
	testDataId       string
	testDataTr       string
	mockedTrainError error
}

func TestGetTrains(t *testing.T) {
	tests := []TrainsTestCase{
		{
			name: "Get_Trains_200",
			url:  "/v1/trains?departure_city=lviv&arrival_city=kyiv&price=300uah",
			want: 200,
		},
		{
			name:             "Get_Trains_404",
			url:              "/v1/trains?mcok=890",
			want:             404,
			mockedGetTrains:  []models.Train{},
			mockedTrainError: http.ErrLineTooLong,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			models.MockedGetData(tc.mockedGetTrains, tc.mockedTrainError)
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
			name:             "Add_To_Trip_201",
			url:              "/v1/trains",
			want:             201,
			testDataId:       "00000000-0000-0000-0000-000000000001",
			testDataTr:       "00000000-0000-0000-0000-000000000001",
			mockedGetTrains:  []models.Train{},
			mockedTrainError: nil,
		},
		{
			name:             "Add_Trains_400",
			url:              "/v1/trains",
			want:             400,
			testDataId:       "00000000-0000-0000-0000-000000000001",
			testDataTr:       "asdsad",
			mockedGetTrains:  []models.Train{},
			mockedTrainError: nil,
		},
		{
			name:             "Add_Trains_400_2",
			url:              "/v1/trains",
			want:             400,
			testDataId:       "asdasd",
			testDataTr:       "00000000-0000-0000-0000-000000000001",
			mockedGetTrains:  []models.Train{},
			mockedTrainError: nil,
		},
		{
			name:             "Add_Events_400_3",
			url:              "/v1/trains",
			want:             400,
			testDataId:       "00000000-0000-0000-0000-000000000001",
			testDataTr:       "00000000-0000-0000-0000-000000000001",
			mockedTrainError: error(new(http.ProtocolError)),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			data := url.Values{}
			data.Add("train_id", tc.testDataId)
			data.Add("trip_id", tc.testDataTr)
			models.MockedAddToTrip(tc.mockedTrainError)

			rec := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodPost, tc.url, bytes.NewBufferString(data.Encode()))
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")

			router.ServeHTTP(rec, req)

			if rec.Code != tc.want {
				t.Errorf("Expected: %d , got %d", tc.want, rec.Code)
			}
			data.Del("train_id")
			data.Del("trip_id")
		})

	}
}

func TestGetTrainFromTrip(t *testing.T) {
	tests := []TrainsTestCase{
		{
			name:            "Get_train_200",
			url:             "/v1/trains/trip/00000000-0000-0000-0000-000000000001",
			want:            200,
			mockedGetTrains: []models.Train{},
		},
		{
			name:            "Get_Trains_400",
			url:             "/v1/trains/trip/sadsad",
			want:            400,
			mockedGetTrains: []models.Train{},
		},
		{
			name:             "Get_Events_404",
			url:              "/v1/trains/trip/00000000-0000-0000-0000-000000000009",
			want:             404,
			mockedGetTrains:  []models.Train{},
			mockedTrainError: http.ErrNoLocation,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			models.MockedGetByTrip(tc.mockedGetTrains, tc.mockedTrainError)
			rec := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodGet, tc.url, nil)

			router.ServeHTTP(rec, req)

			if rec.Code != tc.want {
				t.Errorf("Expected: %d , got %d", tc.want, rec.Code)
			}
		})
	}
}
