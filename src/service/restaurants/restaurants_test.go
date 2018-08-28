package restaurants_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/Nastya-Kruglikova/cool_tasks/src/model"
	"github.com/Nastya-Kruglikova/cool_tasks/src/service"
	"github.com/Nastya-Kruglikova/cool_tasks/src/service/restaurants"
	"net/http"
	"net/http/httptest"
	"testing"
)

var router = service.NewRouter()

type RestaurantsTestCase struct {
	name                 string
	url                  string
	want                 int
	mockedGetRestaurants []model.Restaurant
	testDataId           string
	testDataEv           string
	mockedRestaurantsErr error
}

func TestGetByRequestHandler(t *testing.T) {
	tests := []RestaurantsTestCase{
		{
			name:                 "Get_Restaurants_200",
			url:                  "/v1/restaurants",
			want:                 200,
			mockedGetRestaurants: []model.Restaurant{},
			mockedRestaurantsErr: nil,
		},
		{
			name:                 "Get_Restaurant_404",
			url:                  "/v1/restaurants?mock=890",
			want:                 404,
			mockedGetRestaurants: []model.Restaurant{},
			mockedRestaurantsErr: http.ErrBodyNotAllowed,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			model.MockedGetData(tc.mockedGetRestaurants, tc.mockedRestaurantsErr)
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
	tests := []RestaurantsTestCase{
		{
			name:                 "Add_Restaurants_201",
			url:                  "/v1/restaurants",
			want:                 201,
			testDataId:           "00000000-0000-0000-0000-000000000001",
			testDataEv:           "00000000-0000-0000-0000-000000000001",
			mockedRestaurantsErr: nil,
		},
		{
			name:                 "Add_Restaurants_400",
			url:                  "/v1/restaurants",
			want:                 400,
			testDataId:           "00000000-0000-0000-0000-000000000001",
			testDataEv:           "asdsad",
			mockedRestaurantsErr: nil,
		},
		{
			name:                 "Add_Restaurants_400_2",
			url:                  "/v1/restaurants",
			want:                 400,
			testDataId:           "asdasd",
			testDataEv:           "00000000-0000-0000-0000-000000000001",
			mockedRestaurantsErr: nil,
		},
		{
			name:                 "Add_Restaurants_400_3",
			url:                  "/v1/restaurants",
			want:                 400,
			testDataId:           "00000000-0000-0000-0000-000000000001",
			testDataEv:           "00000000-0000-0000-0000-000000000001",
			mockedRestaurantsErr: error(new(http.ProtocolError)),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			var data restaurants.TripRestaurant
			data.RestaurantID = tc.testDataEv
			data.TripID = tc.testDataId
			body, _ := json.Marshal(data)

			model.MockedAddToTrip(tc.mockedRestaurantsErr)
			rec := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodPost, tc.url, bytes.NewReader(body))
			req.Header.Set("Content-Type", "application/json")
			router.ServeHTTP(rec, req)
			fmt.Println(rec.Code)
			if rec.Code != tc.want {
				t.Errorf("Expected: %d , got %d", tc.want, rec.Code)
			}
		})
	}
}

func TestGetByTripHandler(t *testing.T) {
	tests := []RestaurantsTestCase{
		{
			name:                 "Get_Restaurants_200",
			url:                  "/v1/restaurants/trip/00000000-0000-0000-0000-000000000001",
			want:                 200,
			mockedGetRestaurants: []model.Restaurant{},
			mockedRestaurantsErr: nil,
		},
		{
			name:                 "Get_Restaurants_400",
			url:                  "/v1/restaurants/trip/sadsad",
			want:                 400,
			mockedGetRestaurants: []model.Restaurant{},
			mockedRestaurantsErr: nil,
		},
		{
			name:                 "Get_Restaurants_404",
			url:                  "/v1/restaurants/trip/00000000-0000-0000-0000-000000000009",
			want:                 404,
			mockedGetRestaurants: []model.Restaurant{},
			mockedRestaurantsErr: http.ErrLineTooLong,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			model.MockedGetByTrip(tc.mockedGetRestaurants, tc.mockedRestaurantsErr)
			rec := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodGet, tc.url, nil)

			router.ServeHTTP(rec, req)
			fmt.Println(rec.Code)
			if rec.Code != tc.want {
				t.Errorf("Expected: %d , got %d", tc.want, rec.Code)
			}
		})
	}
}
