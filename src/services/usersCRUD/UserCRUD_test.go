package usersCRUD_test

import (
	"testing"
	"net/http/httptest"
	"net/http"
	"github.com/Nastya-Kruglikova/cool_tasks/src/services"
)

var router = services.NewRouter()

type getUsersTestCase struct {
	name string
	url string
	want int
}

func TestGetUsers(t *testing.T) {
	tests := []getUsersTestCase{
		{
			name: "Get_Users_200",
			url:  "/v1/users",
			want: 200,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodGet, tc.url, nil)

			router.ServeHTTP(rec, req)

			if rec.Code != tc.want {
				t.Errorf("Expected: %d , got %d", tc.want, rec.Code)
			}
		})
	}
}



