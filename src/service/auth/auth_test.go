package auth

import (
	"bytes"
	"github.com/Nastya-Kruglikova/cool_tasks/src/model"
	"github.com/alicebob/miniredis"
	"github.com/satori/go.uuid"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

type authTestCase struct {
	name string
	url  string
	want int
}

func TestLogin(t *testing.T) {
	tests := []authTestCase{
		{
			name: "Login_200",
			url:  "/v1/login",
			want: 200,
		},
	}

	expetedLogin := "admin"
	expetedPass := "admin"

	data := url.Values{}
	data.Add("login", expetedLogin)
	data.Add("password", expetedPass)

	model.GetUserByLogin = func(login string) (model.User, error) {
		UserID, _ := uuid.FromString("00000000-0000-0000-0000-000000000001")

		expected := model.User{
			ID:       UserID,
			Name:     "John",
			Login:    expetedLogin,
			Password: expetedPass,
		}

		return expected, nil
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodPost, tc.url, bytes.NewBufferString(data.Encode()))
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
			MockedLogin(rec, req)

			if rec.Code != tc.want {
				t.Errorf("Expected: %d , got %d", tc.want, rec.Code)
			}
		})
	}

}

func TestLogout(t *testing.T) {
	var redis *miniredis.Miniredis
	redis, _ = miniredis.Run()
	redis.Push("00000000-0000-0000-0000-000000000001", "admin")

	tests := []authTestCase{
		{
			name: "Logout_200",
			url:  "/v1/logout",
			want: 200,
		},
	}
	data := url.Values{}
	data.Add("sessionID", "00000000-0000-0000-0000-000000000001")

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodPost, tc.url, bytes.NewBufferString(data.Encode()))
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")

			MockedLogout(rec, req)

			if rec.Code != tc.want {
				t.Errorf("Expected: %d , got %d", tc.want, rec.Code)
			}
		})
	}
}
