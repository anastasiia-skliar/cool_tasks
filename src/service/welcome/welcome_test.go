package welcome_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Nastya-Kruglikova/cool_tasks/src/service"
)

var router = service.NewRouter()

type getWelcomePageTestCase struct {
	name string
	url  string
	want int
}

func TestGetWelcomeHandler(t *testing.T) {
	tests := []getWelcomePageTestCase{
		{
			name: "Get_Welcome_200",
			url:  "/v1/hello-world",
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
