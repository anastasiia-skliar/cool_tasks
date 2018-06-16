package tasksCRUD_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/Nastya-Kruglikova/cool_tasks/src/services"
	"net/url"
	"bytes"
	"github.com/satori/go.uuid"
	"github.com/Nastya-Kruglikova/cool_tasks/src/models"
)

var router = services.NewRouter()
var testUUID = "00000000-0000-0000-0000-000000000001"
//var testUUIDbad = "00000000-0000-0000-0000-000000000002"
var testID, _ = uuid.FromString("00000000-0000-0000-0000-000000000011")

type tasksCRUDTestCase struct {
	name             string
	url              string
	want             int
	mockedGetTask    models.Task
	mockedGetTasks   []models.Task
	mockedDeleteTask uuid.UUID
	mockedCreateTask models.Task
	mockedTasksError error
}

func TestGetTasks(t *testing.T) {

	tests := []tasksCRUDTestCase{
		{
			name:             "Get_Tasks_200",
			url:              "/v1/tasks",
			want:             200,
			mockedGetTasks:   []models.Task{},
			mockedTasksError: nil,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			models.MockedGetTasks(tc.mockedGetTasks, tc.mockedTasksError)
			rec := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodGet, tc.url, nil)

			router.ServeHTTP(rec, req)

			if rec.Code != tc.want {
				t.Errorf("Expected: %d , got %d", tc.want, rec.Code)
			}
		})
	}
}

func TestGetTasksByID(t *testing.T) {
	tests := []tasksCRUDTestCase{
		{
			name: "Get_TaskById_200",
			url:  "/v1/tasks/"+ testUUID,
			want: 200,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			models.MockedGetTask(tc.mockedGetTask, tc.mockedTasksError)
			rec := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodGet, tc.url, nil)

			router.ServeHTTP(rec, req)

			if rec.Code != tc.want {
				t.Errorf("Expected: %d , got %d", tc.want, rec.Code)
			}
		})
	}
}

func TestDeleteTasks(t *testing.T) {
	tests := []tasksCRUDTestCase{
		{
			name: "Delete_Task_200",
			url:  "/v1/tasks/" + testUUID,
			want: 200,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			models.MockedDeleteTask(tc.mockedDeleteTask, tc.mockedTasksError)
			rec := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodDelete, tc.url, nil)

			router.ServeHTTP(rec, req)

			if rec.Code != tc.want {
				t.Errorf("Expected: %d , got %d", tc.want, rec.Code)
			}
		})
	}
}

func TestAddTasks(t *testing.T) {
	tests := []tasksCRUDTestCase{
		{
			name: "Add_Task_200",
			url:  "/v1/tasks",
			want: 200,
		},
	}

	data := url.Values{}
	data.Add("user_id", "00000000-0000-0000-0000-000000000011") //bad value
	data.Add("name", "JustUser")
	data.Add("time", "Mon Jan 2 15:04:05 MST 2006")
	data.Add("desc", "Desc of my task")

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			models.MockedCreateTask(tc.mockedCreateTask, tc.mockedTasksError)
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
