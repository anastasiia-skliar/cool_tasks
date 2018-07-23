package usersCRUD_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/Nastya-Kruglikova/cool_tasks/src/models"
	"github.com/Nastya-Kruglikova/cool_tasks/src/services"

	"github.com/satori/go.uuid"
)

var router = services.NewRouter()

type usersCRUDTestCase struct {
	name              string
	url               string
	want              int
	mockedGetUser     models.User
	mockedCreateUser  models.User
	mockedGetUsers    []models.User
	mockedUserError   error
	mockedDeleteUsers uuid.UUID
}

func TestGetUsers(t *testing.T) {
	tests := []usersCRUDTestCase{
		{
			name:            "Get_Users_200",
			url:             "/v1/users",
			want:            200,
			mockedGetUsers:  []models.User{},
			mockedUserError: nil,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			models.MockedGetUsers(tc.mockedGetUsers, tc.mockedUserError)
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
	tests := []usersCRUDTestCase{
		{
			name:            "Get_Users_200",
			url:             "/v1/users/a7264252-6ef4-11e8-9982-0242ac110002",
			want:            200,
			mockedGetUser:   models.User{},
			mockedUserError: nil,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			models.MockedGetUser(tc.mockedGetUser, tc.mockedUserError)
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
	userId, _ := uuid.FromString("00000000-0000-0000-0000-000000000001")
	tests := []usersCRUDTestCase{
		{
			name:              "Delete_Users_200",
			url:               "/v1/users/00000000-0000-0000-0000-000000000001",
			want:              200,
			mockedDeleteUsers: userId,
			mockedUserError:   nil,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			models.MockedDeleteUser(userId, nil)
			rec := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodDelete, tc.url, nil)

			router.ServeHTTP(rec, req)

			if rec.Code != tc.want {
				t.Errorf("Expected: %d , got %d", tc.want, rec.Code)
			}
		})
	}
}

func TestCreateUser(t *testing.T) {
	tests := []usersCRUDTestCase{
		{
			name:             "Add_Users_200",
			url:              "/v1/users",
			want:             200,
			mockedCreateUser: models.User{},
			mockedUserError:  nil,
		},
	}
	data := url.Values{}
	data.Add("name", "Karim")
	data.Add("login", "Karim123")
	data.Add("password", "1324qwer")
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			models.MockedCreateUser(tc.mockedCreateUser)
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
