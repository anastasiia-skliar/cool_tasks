package restaurants_test

import (
	"github.com/Nastya-Kruglikova/cool_tasks/src/models"
	"testing"
	"net/http/httptest"
	"net/http"
	"github.com/Nastya-Kruglikova/cool_tasks/src/services"
)

var router = services.NewRouter()

type RestaurantsTestCase struct {
	name             string
	url              string
	want             int
	mockedGetRestaurants []models.Restaurant
	testDataId       string
	testDataRe       string
	mockedRestaurantErr  error
}

func TestGetRestaurants(t *testing.T) {
	tests:=[]RestaurantsTestCase{
		{
			name: "Get_Restaurants_200",
			url:  "/v1/restaurants?location=Lviv",
			want: 200,
			mockedGetRestaurants:[]models.Restaurant{},
			mockedRestaurantErr:nil,
		},
		{
			name: "Get_Restaurants_404",
			url:  "/v1/restaurants?mock=909",
			want: 404,
			mockedGetRestaurants:[]models.Restaurant{},
			mockedRestaurantErr:http.ErrBodyNotAllowed,
		},
		{
			name: "Get_Restaurants_404_2",
			url:  "/v1/restaurants?id=sdsadsa",
			want: 404,
			mockedGetRestaurants:[]models.Restaurant{},
			mockedRestaurantErr:http.ErrBodyNotAllowed,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			models.MockedGetRestaurants(tc.mockedGetRestaurants, tc.mockedRestaurantErr)
			rec := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodGet, tc.url, nil)

			router.ServeHTTP(rec, req)

			if rec.Code != tc.want {
				t.Errorf("Expected: %d , got %d", tc.want, rec.Code)
			}
		})
	}
}