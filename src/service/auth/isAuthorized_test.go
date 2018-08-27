package auth

import (
	"github.com/Nastya-Kruglikova/cool_tasks/src/service/common"
	"net/http"
	"net/http/httptest"
	"testing"
)

var cookieSession = "6c3a65d23c5f26fc529f6c5ce01a6b31"

type isAuthorizedTestCase struct {
	name        string
	url         string
	cookieName  string
	cookieValue string
	want        int
}

func TestIsAuthorized(t *testing.T) {
	tests := []isAuthorizedTestCase{
		{
			name:        "isAuthorized_200",
			url:         "/v1/tasks",
			want:        200,
			cookieName:  "user_session",
			cookieValue: cookieSession,
		},
		{
			name:        "isAuthorized_200",
			url:         "/v1/tasks",
			want:        400,
			cookieName:  "badname",
			cookieValue: "",
		},
		{
			name:        "isAuthorized_200",
			url:         "/v1/tasks",
			want:        401,
			cookieName:  "user_session",
			cookieValue: "badValue",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			mockedIsExistRedis()

			rec := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodGet, tc.url, nil)

			cookie := http.Cookie{Name: tc.cookieName, Value: tc.cookieValue}
			req.AddCookie(&cookie)

			fackedNext := func(w http.ResponseWriter, r *http.Request) {
				common.RenderJSON(w, r, "200")
			}

			IsAuthorized(rec, req, fackedNext)

			if rec.Code != tc.want {
				t.Errorf("Expected: %d , got %d", tc.want, rec.Code)
			}
		})
	}
}
