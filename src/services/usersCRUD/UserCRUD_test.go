package usersCRUD_test

import (
	"testing"
	"net/http/httptest"
	"net/http"
	"github.com/Nastya-Kruglikova/cool_tasks/src/services"
	"bytes"
	"net/url"
)

var router = services.NewRouter()

type getUsersTestCase struct {
	name string
	url  string
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

func TestGetUserByID(t *testing.T) {
	tests := []getUsersTestCase{
		{
			name: "Get_Users_200",
			url:  "/v1/users/00000000-0000-0000-0000-000000000001",
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

func TestDeleteUser(t *testing.T) {
	tests := []getUsersTestCase{
		{
			name: "Delete_Users_200",
			url:  "/v1/users/00000000-0000-0000-0000-000000000001",
			want: 200,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodDelete, tc.url, nil)

			router.ServeHTTP(rec, req)

			if rec.Code != tc.want {
				t.Errorf("Expected: %d , got %d", tc.want, rec.Code)
			}
		})
	}
}

func TestAddUser(t *testing.T) {
	tests := []getUsersTestCase{
		{
			name: "Add_Users_200",
			url:  "/v1/users",
			want: 200,
		},
	}
	data := url.Values{}
	data.Add("name", "Karim")
	data.Add("login", "Karim123")
	data.Add("password", "1324qwer")
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodPost, tc.url, bytes.NewBufferString(data.Encode()))
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")

			router.ServeHTTP(rec, req)

			if rec.Code != tc.want {
				t.Errorf("Expected: %d , got %d", tc.want, rec.Code)
			}
		})
	}
}
